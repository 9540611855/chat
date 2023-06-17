package user

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	"pipiChat/models"
	util "pipiChat/pkg/util"

	"github.com/gin-gonic/gin"
)

type ForgePasswordJson struct {
	Email            string `json:"email" form:"email" xml:"email" binding:"required"`
	VerificationCode string `form:"verification_code" json:"verification_code" xml:"verification_code" binding:"required"`
	Password         string `form:"password" json:"password" xml:"password" binding:"required"`
}

func ForgePasswordHandle(c *gin.Context) {
	var forgePasswordJson ForgePasswordJson
	if err := c.ShouldBindJSON(&forgePasswordJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//1.验证邮箱验证码
	var userEmailRegistration auth.UserEmailRegistration
	userEmailRegistration.Email = forgePasswordJson.Email
	userEmailRegistration.VerificationCode = forgePasswordJson.VerificationCode
	if !auth.ValueCode(userEmailRegistration) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱验证码错误"})
		return
	}
	//2.验证password是否包含特殊字符
	if !util.ContainsSqlInjectionCheck(forgePasswordJson.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码里面存在危险字符"})
		return
	}
	//3.修改密码到数据库里面
	var selectUser models.User

	result := configs.DB.Find(&selectUser, "mail = ?", forgePasswordJson.Email)

	if result.RowsAffected == 0 || result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱不存在"})
		return
	}
	selectUser.Password = forgePasswordJson.Password
	configs.DB.Save(selectUser)
	c.JSON(http.StatusOK, gin.H{"success": "修改密码成功"})
}
