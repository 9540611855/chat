package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	init_ "pipiChat/init"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCheckUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	init_.InitDB()
	// 创建一个模拟的 HTTP 请求
	data := login{
		UserName: "pipixia4",
		Password: "pass6123..",
		Token:    "success",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	req, err := http.NewRequest("POST", "/user/login", bytes.NewReader(jsonData))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	// 创建一个 httptest.ResponseRecorder 以记录响应
	req.Header.Add("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()

	// 创建一个临时的 Gin 引擎并注册处理函数
	router := gin.Default()

	router.POST("/user/login", UserLoginHandle)

	// 将模拟的请求传递给处理函数
	router.ServeHTTP(respRecorder, req)

	// 检查响应状态码和内容
	assert.Equal(t, http.StatusOK, respRecorder.Code)
	t.Log(respRecorder.Body.String())
	assert.Contains(t, respRecorder.Body.String(), "eyJ")
}
