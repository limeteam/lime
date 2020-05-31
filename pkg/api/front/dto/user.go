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
	Uid               int        `json:"uid"`
	Nickname          string     `json:"nickname"`
	Mobile            string     `json:"mobile"`
	UserToken         string     `json:"user_token"`
	IsVip             int        `json:"is_vip"`
	Gender            int        `json:"gender"`
	Avatar            string     `json:"avatar"`
	Remain            string     `json:"remain"`
	GoldRemain        string     `json:"gold_remain"`
	SilverRemain      string     `json:"silver_remain"`
	CoinToday         int        `json:"coin_today"`
	InviteCode        string     `json:"invite_code"`
	TodayReadDuration int        `json:"today_read_duration"`
	IsNewUser         int        `json:"is_new_user"`
	CoinTipTitle      string     `json:"coin_tip_title"`
	MessageNoreadNum  int        `json:"message_noread_num"`
	Amount            int        `json:"amount"`
	Coin              int        `json:"coin"`
	Level             int        `json:"level"`
	SignDays          int        `json:"sign_days"`
	AutoSub           string     `json:"auto_sub"`
	BindList          []BindList `json:"bind_list"`
	Status            int        `json:"status"`
	CreatedAt         string     `json:"created_at"`
}

type BindList struct {
	Label   string `json:"label"`
	Action  string `json:"action"`
	Status  int    `json:"status"`
	Display string `json:"display"`
}
