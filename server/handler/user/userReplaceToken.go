package user

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type repalceTokenJson struct {
	Email            string `json:"email"`
	VerificationCode string `form:"password" json:"password" xml:"password" binding:"required"`
	Token            string `form:"token" json:"token" xml:"token" binding:"required"`
}

func RepalceToken(r repalceTokenJson) bool {
	//1.根据邮箱查询userid
	var selectUser model.User
	result := configs.DB.Find(&selectUser, "mail = ?", r.Email)
	if result.RowsAffected == 0 || result.Error != nil {
		return false
	}
	var userState model.UserState
	res := configs.DB.Find(&userState, "id = ?", selectUser.ID)
	if res.RowsAffected == 0 || res.Error != nil {
		return false
	}
	userState.Token = r.Token
	configs.DB.Save(userState)
	return true
}

func RepalceTokenHandle(c *gin.Context) {
	var repalceToken repalceTokenJson
	if err := c.ShouldBindJSON(&repalceToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//1.验证邮箱code是否正确
	var userEmailRegistration auth.UserEmailRegistration
	userEmailRegistration.Email = repalceToken.Email
	userEmailRegistration.VerificationCode = repalceToken.VerificationCode
	if !auth.ValueCode(userEmailRegistration) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱验证码错误"})
		return
	}
	//2.替换token
	if !RepalceToken(repalceToken) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "数据出错"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "验证成功"})
}
