package group

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type JoinGroup struct {
	Token     string `form:"token" json:"token" binding:"required"`
	UserId    int64  `form:"userId" json:"userId" binding:"required"`
	GroupId   int64  `form:"groupId" json:"groupId" binding:"required"`
	TypeId    int    `form:"typeId" json:"typeId" binding:"required"`
	GroupName string `form:"groupName" json:"groupName" binding:"required"`
}

func (j JoinGroup) SaveDB(groupUser model.GroupUser) bool {
	if err := configs.DB.Create(&groupUser).Error; err != nil {
		return false
	}
	return true
}

func (j JoinGroup) JoinGroup() bool {
	var count int64
	var groupUser model.GroupUser
	if j.TypeId == 1 {
		configs.DB.Model(&model.GroupUser{}).Where("group_id = ? AND user_id = ?", j.GroupId, j.UserId).Count(&count)
		if count > 0 {
			return false
		}

	}
	if j.TypeId == 2 {
		configs.DB.Model(&model.GroupUser{}).Where("group_name = ? AND user_id = ?", j.GroupName, j.UserId).Count(&count)
		if count > 0 {
			return false
		}
	}
	groupUser.GroupID = j.GroupId
	groupUser.UserId = j.UserId
	groupUser.UserType = 3
	if j.SaveDB(groupUser) {
		return false
	}
	return true
}

func (j JoinGroup) JoinGroupHandler(c *gin.Context) {
	if err := c.ShouldBindJSON(&j); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate JWT token

	if !auth.CheckJwtTokenTime(j.Token, j.UserId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "jwt token 错误"})
		return
	}

	// Check if the user is already in the group
	if !j.JoinGroup() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "加入群组错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "请求已发送，等待管理员同意"})
}
