package friend

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type FriendMessage struct {
	Token    string `form:"token" json:"token" xml:"token" binding:"required"`
	UserId   int64  `form:"userId" json:"userId" xml:"userId" binding:"required"`
	FriendId int64  `form:"friendId" json:"friendId" xml:"friendId" binding:"required"`
	Message  string `form:"friendMessage" json:"friendMessage" xml:"friendMessage" binding:"required"`
}

func (f FriendMessage) findFriend(selectUser *model.FriendGroup) bool {
	//如果是0的话 就认为是使用id寻找朋友

	findRes := configs.DB.Find(selectUser, "friend_id = ? and user_id=? ", f.FriendId, f.UserId)
	if findRes.Error != nil || findRes.RowsAffected == 0 {
		return false
	}

	//如果还不是正常好友
	if selectUser.FriendType != 2 {
		return false
	}
	return true

}

func (f FriendMessage) SendMessageHandle(c *gin.Context) {
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var selectUser model.FriendGroup
	//1.验证jwt是否正确

	if !auth.CheckJwtTokenTime(f.Token, f.UserId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jwt token 错误"})
		return
	}

	//2.在数据库里面寻找朋友
	if !f.findFriend(&selectUser) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "您还没有此好友"})
		return
	}
	//3.存储发送信息
	var friendChat model.ChatInformationFriend
	friendChat.SendId = f.UserId
	friendChat.Message = f.Message
	friendChat.FriendId = f.FriendId
	result := configs.DB.Create(&friendChat)
	if result.RowsAffected == 0 || result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "发送信息成功"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "发送信息失败"})

}
