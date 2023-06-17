package models

type UserState struct {
	UserId     int64  `json:"userId" gorm:"not null;unique" description:"用户id"`
	LoginState bool   `json:"loginState" gorm:"not null" description:"登录状态,默认为false"`
	ID         int64  `json:"id" description:"id"`
	Token      string `json:"token" gorm:"not null"  description:"设备唯一token"`
	JwtToken   string `json:"jwt_token" gorm:"not null"  description:"jwt token"`
}
