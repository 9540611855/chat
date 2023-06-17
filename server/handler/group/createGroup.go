package group

import (
	"math/rand"
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateGroup struct {
	Token     string `form:"token" json:"token" binding:"required"`
	UserId    int64  `form:"adminId" json:"adminId" binding:"required"`
	GroupName string `form:"groupName" json:"groupName" binding:"required"`
}

func generateUniqueID() int64 {
	// Get the current timestamp in milliseconds
	now := time.Now().UnixNano() / int64(time.Millisecond)

	// Generate a random number between 0 and 9999
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(10000)

	// Concatenate the timestamp and the random number to create the ID
	id := now*10000 + randomNumber

	return id
}

func CreateGroupHandler(c *gin.Context) {
	var createGroupReq CreateGroup
	if err := c.ShouldBindJSON(&createGroupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate JWT token

	if !auth.CheckJwtTokenTime(createGroupReq.Token, createGroupReq.UserId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "jwt token 错误"})
		return
	}

	// Create a new group
	group := model.GroupInfo{
		GroupId:   generateUniqueID(),
		AdminId:   createGroupReq.UserId,
		GroupName: createGroupReq.GroupName,

		CreateGroupTime: time.Now().UnixNano(),
	}
	if err := configs.DB.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户组失败"})
		return
	}
	/*
		添加管理员自己在群组里面
	*/
	groupUser := model.GroupUser{
		GroupID:  group.GroupId,
		UserId:   group.AdminId,
		UserType: 2,
	}
	if err := configs.DB.Create(&groupUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户组失败"})
		return
	}
	// Return the new group information
	c.JSON(http.StatusOK, gin.H{
		"groupId":         group.ID,
		"adminId":         group.AdminId,
		"groupName":       group.GroupName,
		"createGroupTime": group.CreateGroupTime,
	})
}
