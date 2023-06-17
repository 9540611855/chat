package group

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSaveChatInformationHandler(t *testing.T) {
	// Initialize request payload
	payload := SaveChatRequest{
		Token:   "valid_token",
		GroupId: 1,
		SendId:  1,
		Message: "Hello World!",
	}
	reqBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	// 创建测试请求
	req, err := http.NewRequest("POST", "/chat", bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatal(err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建测试响应
	w := httptest.NewRecorder()

	// 调用被测试的函数
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	a := SaveChatRequest{}
	a.SaveChatInformationHandler(c)

	// 检查响应状态码和内容
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"success":true}`
	assert.Equal(t, expectedBody, w.Body.String())
}

func TestGetChatInformationHandle(t *testing.T) {
	// Initialize request payload
	payload := GetChatRequest{
		Token:   "valid_token",
		GroupId: 1,
		UserId:  1,
	}

	reqBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	// 创建测试请求
	req, err := http.NewRequest("POST", "/chat", bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatal(err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建测试响应
	w := httptest.NewRecorder()

	// 调用被测试的函数
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	g := GetChatRequest{}
	g.GetChatInformationHandle(c)

	// 检查响应状态码和内容
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"success":true}`
	assert.Equal(t, expectedBody, w.Body.String())
}
