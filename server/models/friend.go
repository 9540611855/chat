package models

type Friend struct {
	ID            int64 `json:"id" description:"id"`
	UserId        int64 `json:"userId" gorm:"not null;unique"  description:"用户id"`
	UserFriendsId int64 `json:"userFriendsId" gorm:"not null" description:"朋友id"`
	FriendsType   int   `json:"friendsType" gorm:"not null" description:"朋友关系类型(拉黑、正常、亲密、删除等)"`
	GroupId       int64 `json:"groupId" gorm:"not null;unique" description:"聊天id(这里简单认为两个好友聊天也是一个groupid)"`
}
