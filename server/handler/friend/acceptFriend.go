package friend

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type AcceptFriend struct {
	Token    string `form:"token" json:"token" xml:"token" binding:"required"`
	UserId   int64  `form:"userId" json:"userId" xml:"userId" binding:"required"`
	FriendId int64  `form:"friendId" json:"friendId" xml:"friendId" binding:"required"`
}

func (a AcceptFriend) AcceptFriendHandle(c *gin.Context) {
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//1.验证jwt是否正确

	if !auth.CheckJwtTokenTime(a.Token, a.UserId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jwt token 错误"})
		return
	}

	//2.在数据库里面查找好友请求
	var friendRequest model.FriendGroup
	result := configs.DB.Where("user_id = ? AND friend_id = ? AND friend_type = ?", a.UserId, a.FriendId, 0).First(&friendRequest)
	if result.RowsAffected == 0 || result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有此好友请求或者请求存在异常"})
		return
	}

	//3.更新好友请求为已接受
	friendRequest.FriendType = 2
	result = configs.DB.Save(&friendRequest)
	if result.RowsAffected == 0 || result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "接受好友请求失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "已接受好友请求，并添加好友关系"})
}
