package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	init_ "pipiChat/init"
	model "pipiChat/models"
	"testing"

	"github.com/gin-gonic/gin"
	c "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestUserSqlInjectionCheck(t *testing.T) {
	c.Convey("UserSqlInjectionCheck should return true", t, func() {

		var user model.User
		user.Password = "q1641240899.."
		user.UserName = "pipixia"
		user.RealName = "皮皮虾"
		c.So(UserSqlInjectionCheck(user), c.ShouldBeTrue)
	})

	c.Convey("UserSqlInjectionCheck should return false", t, func() {

		var user model.User
		user.Password = "q1641240899.."
		user.UserName = "pipixia"
		user.RealName = "皮皮虾11"
		c.So(UserSqlInjectionCheck(user), c.ShouldBeFalse)
	})
	c.Convey("UserSqlInjectionCheck should return false", t, func() {

		var user model.User
		user.Password = "123456789'or1"
		user.UserName = "ggflag"
		user.RealName = "皮皮虾"
		c.So(UserSqlInjectionCheck(user), c.ShouldBeFalse)
	})
}

func TestUserNameCheck(t *testing.T) {
	c.Convey("UserNameCheck should return false", t, func() {
		init_.InitDB()
		c.So(userNameCheck("admin"), c.ShouldBeFalse)
	})
	c.Convey("UserNameCheck should return true", t, func() {
		init_.InitDB()
		c.So(userNameCheck("ad12min"), c.ShouldBeTrue)
	})
}

type RegisterData struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Realname string `json:"realname"`
	Token    string `json:"token"`
	Mail     string `json:"mail"`
}

func TestUserRegister(t *testing.T) {

	gin.SetMode(gin.TestMode)
	init_.InitDB()
	// 创建一个模拟的 HTTP 请求
	data := RegisterData{
		Username: "pipixia2",
		Password: "pass6123.",
		Realname: "皮皮虾",
		Token:    "error",
		Mail:     "954061185@qq.com",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	req, err := http.NewRequest("POST", "/user/register", bytes.NewReader(jsonData))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	// 创建一个 httptest.ResponseRecorder 以记录响应
	req.Header.Add("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()

	// 创建一个临时的 Gin 引擎并注册处理函数
	router := gin.Default()

	router.POST("/user/register", UserRegisterHandle)

	// 将模拟的请求传递给处理函数
	router.ServeHTTP(respRecorder, req)

	// 检查响应状态码和内容
	assert.Equal(t, http.StatusOK, respRecorder.Code)
	t.Log(respRecorder.Body.String())
	assert.Contains(t, respRecorder.Body.String(), "注册成功")

}
