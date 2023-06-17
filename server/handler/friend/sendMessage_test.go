package friend

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	init_ "pipiChat/init"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSendMessageHandle(t *testing.T) {
	init_.InitDB()
	r := gin.New()

	// 添加路由处理函数
	r.POST("/friend/message", FriendMessage{}.SendMessageHandle)

	// 创建测试请求
	w := httptest.NewRecorder()
	reqBody := map[string]interface{}{
		"token":         "test",
		"userId":        4,
		"friendId":      27,
		"friendMessage": "hello",
	}
	reqBytes, _ := json.Marshal(reqBody)
	reqReader := strings.NewReader(string(reqBytes))
	req, _ := http.NewRequest("POST", "/friend/message", reqReader)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	r.ServeHTTP(w, req)

	// 检查响应状态码是否为 200
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status %d; got %d\n", http.StatusOK, w.Code)
	}

	// 检查响应内容是否为 JSON 格式
	contentType := w.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("expected content type to be application/json; got %s\n", contentType)
	}

	// 解码响应内容
	var resp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Errorf("failed to decode response: %v\n", err)
	}

	// 检查相应内容是否符合预期
	expectedMsg := "您还没有此好友"
	if resp["error"] != expectedMsg {
		t.Errorf("expected message %q; got %q\n", expectedMsg, resp["success"])
	}
}
