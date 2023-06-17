package group

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type LeaveGroup struct {
	Token   string `form:"token" json:"token" binding:"required"`
	UserId  int64  `form:"userId" json:"userId" binding:"required"`
	GroupId int64  `form:"groupId" json:"groupId" binding:"required"`
}

func (l LeaveGroup) DeleteGroupUser() bool {
	if err := configs.DB.Where("group_id = ? AND user_id = ?", l.GroupId, l.UserId).Delete(&model.GroupUser{}).Error; err != nil {
		return false
	}
	return true
}

func (l LeaveGroup) LeaveGroup() bool {
	var count int64
	configs.DB.Model(&model.GroupUser{}).Where("group_id = ? AND user_id = ?", l.GroupId, l.UserId).Count(&count)
	if count == 0 {
		return false
	}

	if !l.DeleteGroupUser() {
		return false
	}
	return true
}

func (l LeaveGroup) LeaveGroupHandler(c *gin.Context) {
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate JWT token

	if !auth.CheckJwtTokenTime(l.Token, l.UserId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "jwt token 错误"})
		return
	}

	// Check if the user is already in the group
	if !l.LeaveGroup() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "退群错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "已退出群组"})
}
