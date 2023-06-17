package models

type FriendGroup struct {
	ID              int64 `json:"ID"`
	UserId          int64 `json:"userId" gorm:"not null;unique" description:"用户id"`
	FriendId        int64 `json:"friendId" gorm:"not null;unique" description:"朋友"`
	FriendType      int   `json:"friendType" gorm:"not null" description:"朋友状态(拉黑-1、未通过0、删除1、正常2)"`
	CreateGroupTime int64 `json:"createGroupTime" gorm:"autoUpdateTime:nano" description:"加好友时间"`
}
