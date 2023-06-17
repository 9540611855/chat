package friend

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type FlashMessage struct {
	Token  string `form:"token" json:"token" xml:"token" binding:"required"`
	UserId int64  `form:"userId" json:"userId" xml:"userId" binding:"required"`
	Time   int64  `form:"time" json:"time" xml:"time" binding:"required"`
}

func (f FlashMessage) findMessage(chatInformation *[]model.ChatInformationFriend) bool {
	// 查询在最后更新时间之后的聊天信息

	if err := configs.DB.Where("time > ?", f.Time).Find(&chatInformation).Error; err != nil {
		return false
	}
	return true
}
func (f FlashMessage) FlashFriendHandle(c *gin.Context) {
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var chatInformation []model.ChatInformationFriend

	//1.验证jwt是否正确
	if !auth.CheckJwtTokenTime(f.Token, f.UserId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jwt token 错误"})
		return
	}

	//2.寻找数据
	if !f.findMessage(&chatInformation) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "服务器拉取数据错误"})
		return
	}
	c.JSON(http.StatusOK, chatInformation)

}
