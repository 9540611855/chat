package group

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type KickoutUserRequest struct {
	Token   string `form:"token" json:"token" binding:"required"`
	GroupId int64  `form:"groupId" json:"groupId" binding:"required"`
	UserId  int64  `form:"userId" json:"userId" binding:"required"`
	AdminId int64  `form:"adminId" json:"adminId" binding:"required"`
}

func (k KickoutUserRequest) checkAdmin() bool {
	groupInfo := model.GroupInfo{}

	err := configs.DB.Where("group_id = ? AND admin_id = ?", k.GroupId, k.AdminId).
		First(&groupInfo).Error

	if err != nil {
		return false
	}
	return true
}

func (k KickoutUserRequest) KickoutUserHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&k); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate JWT token
	if !auth.CheckJwtTokenTime(k.Token, k.AdminId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "jwt token 错误"})
		return
	}
	if !k.checkAdmin() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "admin用户信息没有查到"})
		return
	}

	groupUser := model.GroupUser{}

	err := configs.DB.Where("group_id = ? AND user_id = ?", k.GroupId, k.UserId).
		First(&groupUser).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询用户失败"})
		return
	}

	if groupUser.UserType == 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该用户已经是申请用户"})
		return
	}
	if groupUser.UserType == 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该用户已经是管理员"})
		return
	}

	groupUser.UserType = -1

	if err := configs.DB.Save(&groupUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "踢出用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
