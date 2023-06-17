package models

type User struct {
	UserName    string `json:"userName" gorm:"not null;unique" description:"用户名"`
	Password    string `json:"password" gorm:"not null" description:"用户密码"`
	ID          int64  `json:"id" description:"id"`
	RealName    string `json:"realName" gorm:"not null"  description:"真实姓名"`
	Permissions int64  `json:"permissions" gorm:"not null" description:"用户权限"`
	Mail        string `json:"mail" gorm:"not null;unique" description:"用户邮箱"`
}
