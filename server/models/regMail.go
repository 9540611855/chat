package models

type RegMail struct {
	Mail      string `json:"mail" gorm:"not null"  description:"用户邮箱"`
	RegNum    int    `json:"reg_num" gorm:"not null"  description:"注册次数"`
	ID        int64  `json:"id" description:"id"`
	MailState bool   `json:"mail_state" gorm:"not null"  description:"邮箱状态0未验证.1是已验证"`
	UserId    int64  `json:"userId" gorm:"not null;unique" description:"用户id"`
}
