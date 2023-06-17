package user

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"
	util "pipiChat/pkg/util"

	"github.com/gin-gonic/gin"
)

type EditPasswordJson struct {
	JwtToken string `form:"jwt_token" json:"jwt_token" xml:"jwt_token" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func EditPasswordHandle(c *gin.Context) {
	var editPasswordJson EditPasswordJson
	if err := c.ShouldBindJSON(&editPasswordJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//1.验证jwt token
	claims, err := auth.ValidateJWT(editPasswordJson.JwtToken)
	if err != nil || claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//2.验证密码是否包含sql注入
	if !util.ContainsSqlInjectionCheck(editPasswordJson.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码里面存在危险字符"})
		return
	}
	//3.修改密码到数据库里面
	userId := claims["user"]
	var selectUser model.User
	id := int64(userId.(float64))
	result := configs.DB.Find(&selectUser, "id = ?", id)

	if result.RowsAffected == 0 || result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jwttoken错误"})
		return
	}
	selectUser.Password = editPasswordJson.Password
	configs.DB.Save(selectUser)

	c.JSON(http.StatusOK, gin.H{"success": "修改密码成功"})
}
