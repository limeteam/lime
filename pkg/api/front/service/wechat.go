package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	wechatSdk "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/spf13/viper"
	"lime/pkg/api/front/dao"
	"lime/pkg/api/front/domain/wechat"
)
var ConfigDao = dao.ConfigDao{}

type WechatService struct{}
type wechatSettingBase struct {
	Url string
	Token string
	Encodingaeskey string
}
func (ws WechatService) Callback(c *gin.Context) {
	wc := wechatSdk.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:        viper.GetString("redis.host"),
		Database:    viper.GetInt("redis.db"),
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60, //second
	}
	redisCache := cache.NewRedis(redisOpts)
	config := ConfigDao.GetByCode("wechatSetting")
	value := config.Config_value
	settingBase := &wechatSettingBase{}
	errSetting := json.Unmarshal(value,&settingBase)
	if errSetting != nil {
		return
	}
	cfg := &offConfig.Config{
		AppID:          viper.GetString("wechat.appid"),
		AppSecret:      viper.GetString("wechat.appsecret"),
		Token:          settingBase.Token,
		EncodingAESKey: settingBase.Encodingaeskey,
		Cache:          redisCache,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	server := officialAccount.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		switch msg.Event {
		case message.EventSubscribe://关注公众号事件
			return wechat.Subscribe()
		case message.EventUnsubscribe:
			break
		default:
			break
		}
		text := message.NewText("欢迎使用lime soft!")
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		return
	}
	//发送回复的消息
	server.Send()
	return
}
