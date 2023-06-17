package models

type GroupUser struct {
	ID            int64 `json:"id"`
	GroupID       int64 `json:"groupID" gorm:"not null;unique" description:"用户组id"`
	JoinGroupTime int64 `json:"joinGroupTime" gorm:"autoUpdateTime:nano" description:"用户加入聊天组的时间"`
	UserId        int64 `json:"userId" gorm:"not null;unique" description:"用户id"`
	UserType      int64 `json:"userType" gorm:"not null;unique" description:"用户的状态(退群1、被踢出群-1 在申请3 管理员 2 进群1)"`
}
