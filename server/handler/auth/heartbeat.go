package auth

import (
	"net/http"
	"pipiChat/configs"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type heartbeatJson struct {
	UserId   int64  `json:"userid" form:"userid" binding:"required" xml:"userid"`
	Token    string `json:"token" form:"token" binding:"required" xml:"token"`
	JwtToken string `json:"jwttoken" form:"jwttoken" binding:"required" xml:"jwttoken"`
}

func checkJwtToken(heartBeatJson heartbeatJson) bool {

	userId := heartBeatJson.UserId
	jwtToken := heartBeatJson.JwtToken
	token := heartBeatJson.Token
	var selectUserState model.UserState

	//1.先查询一下是否已经存在用户
	findRes := configs.DB.Find(&selectUserState, "user_id = ?", userId)
	if findRes.RowsAffected == 0 || findRes.Error != nil {
		return false
	}
	//2.查询一下jwttoken和设备token是否一致
	if selectUserState.JwtToken != jwtToken ||
		token != selectUserState.Token {
		return false
	}
	//3.校验时间是否超时
	if !CheckJwtTokenTime(jwtToken, userId) {
		selectUserState.LoginState = false
		configs.DB.Save(&selectUserState)
		return false
	}
	return true
}

// 重发jwttoken
func regenerateJwtToken(userId int64) bool {
	var selectUserState model.UserState
	//1.先查询一下是否已经存在用户
	findRes := configs.DB.Find(&selectUserState, "user_id = ?", userId)
	if findRes.RowsAffected == 0 || findRes.Error != nil {
		return false
	}
	jwtToken, err := GenerateJWT(userId)
	if err != nil {
		return false
	}
	selectUserState.JwtToken = jwtToken
	configs.DB.Save(&selectUserState)
	return true
}

func HeartbeatHandler(c *gin.Context) {
	var bindJson heartbeatJson
	if err := c.ShouldBind(&bindJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !checkJwtToken(bindJson) {
		c.JSON(http.StatusOK, gin.H{"error": "jwtToken失效"})
		return
	}
	if !regenerateJwtToken(bindJson.UserId) {
		c.JSON(http.StatusOK, gin.H{"error": "验证jwtoken失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "下发token成功"})
	return
}
