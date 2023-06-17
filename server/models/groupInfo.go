package models

type GroupInfo struct {
	ID              int64  `json:"ID"`
	GroupId         int64  `json:"groupId" gorm:"not null;unique" description:"用户id"`
	AdminId         int64  `json:"adminId" gorm:"not null" description:"管理员id"`
	GroupName       string `json:"groupName" gorm:"not null" description:"聊天组名字"`
	CreateGroupTime int64  `json:"createGroupTime" gorm:"not null" description:"创建用户组时间"`
}
