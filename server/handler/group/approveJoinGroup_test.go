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

func TestApproveJoinRequestHandler(t *testing.T) {
	// 构造请求参数
	reqBody := ApproveJoinRequest{
		Token:   "your_jwt_token_here",
		GroupId: 123,
		UserId:  456,
		AdminId: 789,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	// 创建测试请求
	req, err := http.NewRequest("POST", "/approve_join_request", bytes.NewBuffer(reqBytes))
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
	a := ApproveJoinRequest{}
	a.ApproveJoinRequestHandler(c)

	// 检查响应状态码和内容
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"success":true}`
	assert.Equal(t, expectedBody, w.Body.String())
}
