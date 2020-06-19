package auth

import (
	"fmt"
	"github.com/go-redis/redis"
	"lime/pkg/common/cache"
	"strconv"
	"time"
)

type AuthInterface interface {
	CreateAuth(int, *TokenDetails) error
	FetchAuth(string) (int, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

type ClientData struct {
	client *redis.Client
}

type AccessDetails struct {
	TokenUuid string
	UserId    int
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

//Save token metadata to Redis
func CreateAuth(userid int, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	err := cache.Set(td.TokenUuid, strconv.Itoa(int(userid)), int(at.Sub(now)))
	if err != nil {
		return err
	}
	errs := cache.Set(td.RefreshUuid, strconv.Itoa(int(userid)), int(rt.Sub(now)))
	if errs != nil {
		return errs
	}
	return nil
}

//Check the metadata saved
func FetchAuth(tokenUuid string) (int, error) {
	userid, err := cache.Get(tokenUuid)
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.Atoi(userid)
	return userID, nil
}

//Once a user row in the token table
func DeleteTokens(authD *AccessDetails) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%d", authD.TokenUuid, authD.UserId)
	//delete access token
	err := cache.Del(authD.TokenUuid)
	if err != nil {
		return err
	}
	//delete refresh token
	errs := cache.Del(refreshUuid)
	if errs != nil {
		return errs
	}
	return nil
}

func DeleteRefresh(refreshUuid string) error {
	//delete refresh token
	err := cache.Del(refreshUuid)
	if err != nil {
		return err
	}
	return nil
}
