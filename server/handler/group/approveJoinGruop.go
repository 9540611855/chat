package group

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type ApproveJoinRequest struct {
	Token   string `form:"token" json:"token" binding:"required"`
	GroupId int64  `form:"groupId" json:"groupId" binding:"required"`
	UserId  int64  `form:"userId" json:"userId" binding:"required"`
	AdminId int64  `form:"adminId" json:"adminId" binding:"required"`
}

func (a ApproveJoinRequest) checkAdmin() bool {
	groupInfo := model.GroupInfo{}

	err := configs.DB.Where("group_id = ? AND admin_id = ?", a.GroupId, a.AdminId).
		First(&groupInfo).Error

	if err != nil {
		return false
	}
	return true
}

func (a ApproveJoinRequest) ApproveJoinRequestHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate JWT token
	if !auth.CheckJwtTokenTime(a.Token, a.AdminId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "jwt token 错误"})
		return
	}
	if !a.checkAdmin() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "admin用户信息没有查到"})
		return
	}

	groupUser := model.GroupUser{}

	err := configs.DB.Where("group_id = ? AND user_id = ?", a.GroupId, a.UserId).
		First(&groupUser).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询加群请求失败"})
		return
	}

	if groupUser.UserType != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该用户不是申请用户"})
		return
	}

	groupUser.UserType = 2

	if err := configs.DB.Save(&groupUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "同意加群请求失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
