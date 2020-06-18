package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/twinj/uuid"
	"net/http"
	"strconv"
	"strings"
	"time"
)


type Token struct{}

func NewToken() *Token {
	return &Token{}
}

type TokenInterface interface {
	CreateToken(userid int) (*TokenDetails, error)
	ExtractTokenMetadata(*http.Request) (*AccessDetails, error)
}

//Token implements the TokenInterface
var _ TokenInterface = &Token{}

func (t *Token) CreateToken(userid int) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.TokenUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + strconv.Itoa(int(userid))

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(viper.GetString("auth.access_secret")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(viper.GetString("auth.refresh_secret")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("auth.access_secret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get the token from the request body
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (t *Token) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	fmt.Println("WE ENTERED METADATA")
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			accessUuid, ok := claims["access_uuid"].(string)
			if !ok {
				return nil, err
			}
			userId, err := strconv.Atoi(fmt.Sprintf("%.f", claims["user_id"]))
			if err != nil {
				return nil, err
			}
			return &AccessDetails{
				TokenUuid: accessUuid,
				UserId:    userId,
			}, nil
	}
	return nil, err
}
