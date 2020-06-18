package dto

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterDto struct {
	Username    string `json:"username"`
	Mobile      string `json:"mobile"`
	Sex         int    `json:"sex"`
	Password    string `json:"password"`
	Salt        string `json:"salt"`
	Faceicon    string `json:"faceicon"`
	Wechat      string `json:"wechat"`
	Email       string `json:"email"`
	Amount      int    `json:"amount"`
	Coin        int    `json:"coin"`
	ExemptLogin int    `json:"exempt_login"`
	Source      int    `json:"source"`
	IsVip       int    `json:"is_vip"`
	ChannelId   int    `json:"channel_id"`
}

type LoginInfoDto struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpires int `json:"at_expires"`
	RtExpires int `json:"rt_expires"`
}

type BindList struct {
	Label   string `json:"label"`
	Action  string `json:"action"`
	Status  int    `json:"status"`
	Display string `json:"display"`
}
