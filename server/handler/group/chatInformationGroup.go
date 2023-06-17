package group

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"
	"time"

	"github.com/gin-gonic/gin"
)

type SaveChatRequest struct {
	Token   string `form:"token" json:"token" binding:"required"`
	GroupId int64  `form:"groupId" json:"groupId" binding:"required"`
	SendId  int64  `form:"sendId" json:"sendId" binding:"required"`
	Message string `form:"message" json:"message" binding:"required"`
}

// 维护群内聊天
func (s SaveChatRequest) SaveChatInformation() error {
	chatInfo := model.ChatInformation{
		SendId:  s.SendId,
		Message: s.Message,
		Time:    time.Now().UnixNano(),
		GroupId: s.GroupId,
	}
	if err := configs.DB.Create(&chatInfo).Error; err != nil {
		return err
	}
	return nil
}

func (s SaveChatRequest) SaveChatInformationHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate JWT token
	if !auth.CheckJwtTokenTime(s.Token, s.SendId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "jwt token 错误"})
		return
	}
	groupUser := model.GroupUser{}

	err := configs.DB.Where("group_id = ? AND user_id = ?", s.GroupId, s.SendId).
		First(&groupUser).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "发送信息失败"})
		return
	}
	if s.SaveChatInformation() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "发送信息失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "发送信息成功"})

}

type GetChatRequest struct {
	Token   string `form:"token" json:"token" binding:"required"`
	GroupId int64  `form:"groupId" json:"groupId" binding:"required"`
	UserId  int64  `form:"sendId" json:"sendId" binding:"required"`
}

// 获取群内聊天记录
func (g GetChatRequest) GetChatInformation() ([]model.ChatInformation, error) {
	chatInfo := []model.ChatInformation{}
	err := configs.DB.Where("group_id = ?", g.GroupId).Find(&chatInfo).Limit(50).Error
	if err != nil {
		return nil, err
	}
	return chatInfo, nil
}

func (g GetChatRequest) GetChatInformationHandle(c *gin.Context) {

	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate JWT token
	if !auth.CheckJwtTokenTime(g.Token, g.UserId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "jwt token 错误"})
		return
	}
	groupUser := model.GroupUser{}

	err := configs.DB.Where("group_id = ? AND user_id = ?", g.GroupId, g.UserId).
		First(&groupUser).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "发送信息失败"})
		return
	}
	ChatInfo, err := g.GetChatInformation()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "发送信息失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"chat": ChatInfo})

}
