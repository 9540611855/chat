package models

type ChatInformationFriend struct {
	ID       int64  `json:"ID"`
	SendId   int64  `json:"sendId"  gorm:"not null" description:"发送信息的用户id"`
	Message  string `json:"message"  gorm:"not null" description:"接收到的信息"`
	Time     int64  `json:"time"  gorm:"autoUpdateTime:nano" description:"接收到的信息时间"`
	FriendId int64  `json:"friendId"  gorm:"not null" description:"信息发送的用户组"`
}
