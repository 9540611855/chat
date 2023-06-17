package friend

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type BlockFriend struct {
	Token    string `form:"token" json:"token" xml:"token" binding:"required"`
	UserId   int64  `form:"userId" json:"userId" xml:"userId" binding:"required"`
	FriendId int64  `form:"friendId" json:"friendId" xml:"friendId" binding:"required"`
}

func (b BlockFriend) findFriend(selectUser *model.FriendGroup) bool {
	//如果是0的话 就认为是使用id寻找朋友

	findRes := configs.DB.Find(selectUser, "friend_id = ? and user_id=? ", b.FriendId, b.UserId)
	if findRes.Error != nil || findRes.RowsAffected == 0 {
		return false
	}

	//如果还不是正常好友
	if selectUser.FriendType != 2 {
		return false
	}
	return true

}

func (b BlockFriend) BlockFriendHandle(c *gin.Context) {
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var selectUser model.FriendGroup

	//1.验证jwt是否正确
	if !auth.CheckJwtTokenTime(b.Token, b.UserId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jwt token 错误"})
		return
	}

	//2.在数据库里面寻找朋友
	if !b.findFriend(&selectUser) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "您还没有此好友"})
		return
	}
	//3.拉黑好友
	var friendRequest model.FriendGroup
	friendRequest.UserId = b.UserId
	friendRequest.FriendType = -1
	friendRequest.FriendId = b.FriendId
	friendRequest.ID = selectUser.ID
	result := configs.DB.Save(&friendRequest)
	if result.RowsAffected == 0 || result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "拉黑好友失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "拉黑好友成功"})

}
