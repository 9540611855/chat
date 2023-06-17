package user

import (
	"net/http"
	"pipiChat/configs"
	auth "pipiChat/handler/auth"
	model "pipiChat/models"
	util "pipiChat/pkg/util"

	"github.com/gin-gonic/gin"
)

type login struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Token    string `form:"token" json:"token" xml:"token" binding:"required"`
}

func checkSqlInject(userLogin login) bool {
	if !util.ContainsSqlInjectionCheck(userLogin.UserName) ||
		!util.ContainsSqlInjectionCheck(userLogin.Password) {
		return false
	}
	return true
}

func checkUser(userLogin login) (bool, int64) {
	//检查用户名是否存在
	var selectUser model.User
	result := configs.DB.Find(&selectUser, "user_name = ?", userLogin.UserName)
	if result.RowsAffected == 0 {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "注册失败,此用户名已被注册"})
		return false, 0
	}
	if result.Error != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return false, 0
	}
	if selectUser.Password == userLogin.Password {
		return true, selectUser.ID
	}
	return false, 0
}

func saveUserState(userLogin login, userId int64, c *gin.Context) bool {
	//result := configs.DB.Create(&userLogin)
	//3.生成jwttoken

	jwtToken, err := auth.GenerateJWT(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "系统错误"})
		return false
	}

	var selectUser model.User
	var selectUserState model.UserState
	result := configs.DB.Find(&selectUser, "user_name = ?", userLogin.UserName)

	if result.RowsAffected == 0 || result.Error != nil {
		return false
	}

	//1.先查询一下是否已经存在用户
	findRes := configs.DB.Find(&selectUserState, "user_id = ?", selectUser.ID)
	if findRes.Error != nil || findRes.RowsAffected == 0 {
		return false
	}
	/*
		//2.不为空就判断一下token是否一致
		if selectUserState.Token != userLogin.Token {
			//todo 进行邮箱验证

			c.JSON(http.StatusBadRequest, gin.H{"error": "检测到设备更换,请验证邮箱"})
			return true
		}
	*/
	//3.存储状态表的token
	selectUserState.JwtToken = jwtToken
	/*
		//4.如果是已经在登录状态 直接返回
		if selectUserState.LoginState {
			c.JSON(http.StatusOK, gin.H{"error": "用户已经在其他设备登录"})
			return true
		}
	*/
	//5.设置在登录状态
	selectUserState.LoginState = true
	configs.DB.Save(selectUserState)
	//6.返回jwt token 以及个人信息
	c.JSON(http.StatusOK, gin.H{"name": selectUser.UserName, "id": selectUser.ID, "email": selectUser.Mail, "jwt": jwtToken})
	//c.JSON(http.StatusOK, gin.H{"success": jwtToken})
	return true
}

func UserLoginHandle(c *gin.Context) {
	var loginJson login
	if err := c.ShouldBindJSON(&loginJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//todo:token鉴权、在数据库里面查找用户密码、防止sql注入等问题 jwt token
	//1.验证账户名密码是否有sql注入
	if !checkSqlInject(loginJson) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "有非法字符"})
		return
	}
	//2.验证账号密码是否正确
	checkRes, userId := checkUser(loginJson)
	if !checkRes {
		c.JSON(http.StatusBadRequest, gin.H{"error": "账号或者密码不对"})
		return
	}
	//4.在用户状态表存下是登录状态
	res := saveUserState(loginJson, userId, c)
	if !res {
		c.JSON(http.StatusBadRequest, gin.H{"error": "系统原因用户登录错误"})
		return
	}
}
