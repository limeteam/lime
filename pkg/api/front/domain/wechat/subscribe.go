package wechat

import "github.com/silenceper/wechat/v2/officialaccount/message"

func Subscribe() *message.Reply {
	text := message.NewText("欢迎关注lime soft!")
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}
