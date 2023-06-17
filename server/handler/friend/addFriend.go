package friend

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddFriend struct {
	Token      string `form:"token" json:"token" xml:"token" binding:"required"`
	UserId     int64  `form:"user_id" json:"user_id" xml:"user_id" binding:"required"`
	FriendId   int64  `form:"friend_id" json:"friend_id" xml:"friend_id" binding:"required"`
	FriendName string `form:"friend_name" json:"friend_name" xml:"friend_name" binding:"required"`
	AddType    int    `form:"add_type" json:"add_type" xml:"add_type" binding:"required"`
}

func (a AddFriend) findFriend(selectUser *model.User) bool {
	//如果是2的话 就认为是使用id寻找朋友
	var findRes *gorm.DB
	if a.AddType == 2 {
		findRes = configs.DB.Find(selectUser, "id = ?", a.FriendId)
		if findRes.Error != nil || findRes.RowsAffected == 0 {
			return false
		}

	}
	//如果是1的话 就认为使用名字寻找朋友
	if a.AddType == 1 {
		findRes = configs.DB.Find(selectUser, "user_name = ?", a.FriendName)
		if findRes.Error != nil || findRes.RowsAffected == 0 {
			return false
		}
	}
	//如果朋友的id和自己的是一样的 那就说明在查找自己
	if selectUser.ID == a.UserId {
		return false
	}

	return true

}

func (a AddFriend) AddFriendHandle(c *gin.Context) {
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var selectUser model.User
	//1.验证jwt是否正确

	if !auth.CheckJwtTokenTime(a.Token, a.UserId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jwt token 错误"})
		return
	}

	//2.在数据库里面寻找朋友
	if !a.findFriend(&selectUser) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有此用户或者用户存在异常"})
		return
	}
	//3.存储加好友请求
	var friendRequest model.FriendGroup
	friendRequest.UserId = a.UserId
	friendRequest.FriendType = 0
	friendRequest.FriendId = a.FriendId
	result := configs.DB.Create(&friendRequest)
	if result.RowsAffected == 0 || result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "发送加好友请求失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "发送加好友请求成功"})

}
