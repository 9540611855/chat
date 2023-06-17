package user

import (
	"net/http"
	"pipiChat/configs"
	model "pipiChat/models"
	util "pipiChat/pkg/util"

	"github.com/gin-gonic/gin"
)

type register struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	RealName string `form:"realname" json:"realname" xml:"realname" binding:"required"`
	Mail     string `form:"mail" json:"mail" xml:"mail" binding:"required"`
	Token    string `form:"token" json:"token" xml:"token" binding:"required"`
}

func userNameCheck(userName string) (checkRes bool) {
	var selectUser model.User

	//检查用户名是否存在
	result := configs.DB.Find(&selectUser, "user_name = ?", userName)
	if result.RowsAffected != 0 {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "注册失败,此用户名已被注册"})
		return false
	}
	if result.Error != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return false
	}
	return true
}

func UserSqlInjectionCheck(user model.User) (checkRes bool) {
	if !util.ContainsSqlInjectionCheck(user.UserName) ||
		!util.ContainsSqlInjectionCheck(user.Password) ||
		!util.RealNameCheck(user.RealName) ||
		!util.ContainsSqlInjectionCheck(user.Mail) {
		return false
	}

	return true

}

func checkRegMail(Mail string) bool {
	var regMail model.RegMail
	result := configs.DB.Find(&regMail, "mail = ?", Mail)

	if result.RowsAffected != 0 || result.Error != nil {
		return false
	}
	return true
}

func registerCheck(c *gin.Context, user model.User) bool {
	//1.检测是否有sql注入的风险
	checkRes := UserSqlInjectionCheck(user)
	if !checkRes {
		c.JSON(http.StatusBadRequest, gin.H{"error": "注册字符有特殊字符"})
		return false
	}
	//2.检测是否为弱密码
	checkRes = util.WeakPasswordCheck(user.Password)
	if !checkRes {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码强度过低"})
		return false
	}
	//3.检测用户名是否已经被注册
	checkRes = userNameCheck(user.UserName)
	if !checkRes {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已经被注册"})
		return false
	}
	//4.检测邮箱是否已经被注册
	if !checkRegMail(user.Mail) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱已经被注册"})
		return false
	}

	return true
}

func UserRegisterHandle(c *gin.Context) {
	var registerJson register
	if err := c.ShouldBindJSON(&registerJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.User
	user.Password = registerJson.Password
	user.RealName = registerJson.RealName
	user.UserName = registerJson.UserName
	user.Permissions = 1
	user.Mail = registerJson.Mail
	if !registerCheck(c, user) {
		return
	}
	//将用户写入user 表
	result := configs.DB.Create(&user)
	if result.RowsAffected == 0 || result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "系统发生了问题，请联系管理员"})
		return
	}
	//将用户写入 用户状态表
	var userState model.UserState

	userState.LoginState = false
	userState.UserId = user.ID
	userState.Token = registerJson.Token

	res := configs.DB.Create(&userState)

	if res.RowsAffected == 0 || res.Error != nil {
		//将存入的数据删除了
		configs.DB.Delete(&user)
		c.JSON(http.StatusBadRequest, gin.H{"error": "系统发生了问题，请联系管理员"})
		return
	}

	//写入邮箱注册表
	var regMail model.RegMail
	regMail.Mail = registerJson.Mail
	regMail.MailState = false
	regMail.RegNum = 1
	regMail.UserId = user.ID
	resRegMail := configs.DB.Create(&regMail)
	if resRegMail.RowsAffected == 0 || resRegMail.Error != nil {
		//将存入的数据删除了
		configs.DB.Delete(&user)
		configs.DB.Delete(&userState)
		c.JSON(http.StatusBadRequest, gin.H{"error": "系统发生了问题，请联系管理员"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "注册成功"})
	return
}
