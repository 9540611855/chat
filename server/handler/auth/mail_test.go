package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValueCodeHandle(t *testing.T) {
	router := gin.Default()
	router.POST("/valuecode", ValueCodeHandle)
	MailstoreArray = append(MailstoreArray, UserEmailRegistration{"test1@example.com", "123456"})
	MailstoreArray = append(MailstoreArray, UserEmailRegistration{"test2@example.com", "654321"})
	// Test case 1: email not registered
	req1Body := `{"email": "test1@example.com", "verification_code": "123456"}`
	req1, _ := http.NewRequest("POST", "/valuecode", strings.NewReader(req1Body))
	req1.Header.Set("Content-Type", "application/json")
	resp1 := httptest.NewRecorder()
	router.ServeHTTP(resp1, req1)
	assert.Equal(t, http.StatusOK, resp1.Code)
	assert.Contains(t, resp1.Body.String(), `"code":1`)
	assert.Contains(t, resp1.Body.String(), `"msg":"ok"`)

	// Test case 2: email registered, wrong verification code
	req2Body := `{"email": "test2@example.com", "verification_code": "123456"}`
	req2, _ := http.NewRequest("POST", "/valuecode", strings.NewReader(req2Body))
	req2.Header.Set("Content-Type", "application/json")
	resp2 := httptest.NewRecorder()
	router.ServeHTTP(resp2, req2)
	assert.Equal(t, http.StatusOK, resp2.Code)
	assert.Contains(t, resp2.Body.String(), `"code":0`)
	assert.Contains(t, resp2.Body.String(), `"msg":"验证码错误"`)

	// Test case 3: email registered, correct verification code
	req3Body := `{"email": "test3@example.com", "verification_code": "123456"}`
	req3, _ := http.NewRequest("POST", "/valuecode", strings.NewReader(req3Body))
	req3.Header.Set("Content-Type", "application/json")
	resp3 := httptest.NewRecorder()
	router.ServeHTTP(resp3, req3)
	assert.Equal(t, http.StatusOK, resp3.Code)
	assert.Contains(t, resp3.Body.String(), `"code":0`)
	assert.Contains(t, resp3.Body.String(), `"msg":"此邮箱并未注册"`)
	// Test case 4: email registered, correct verification code
	req4Body := `{"email": "test1@example.com", "verification_code": "123456"}`
	req4, _ := http.NewRequest("POST", "/valuecode", strings.NewReader(req4Body))
	req4.Header.Set("Content-Type", "application/json")
	resp4 := httptest.NewRecorder()
	router.ServeHTTP(resp4, req4)
	assert.Equal(t, http.StatusOK, resp4.Code)
	assert.Contains(t, resp4.Body.String(), `"code":0`)
	assert.Contains(t, resp4.Body.String(), `"msg":"此邮箱并未注册"`)
}

func TestSendMailHandle(t *testing.T) {
	router := gin.Default()
	router.POST("/sendmail", SendMailHandle)

	// Test case 1: invalid request body
	req1Body := `{"email: "test1@example.com"}`
	req1, _ := http.NewRequest("POST", "/sendmail", strings.NewReader(req1Body))
	req1.Header.Set("Content-Type", "application/json")
	resp1 := httptest.NewRecorder()
	router.ServeHTTP(resp1, req1)
	assert.Equal(t, http.StatusOK, resp1.Code)
	assert.Contains(t, resp1.Body.String(), `"code":0`)
	assert.Contains(t, resp1.Body.String(), `"msg":"error"`)

	// Test case 2: email sent successfully
	req2Body := `{"email": "m18568675386@163.com"}`
	req2, _ := http.NewRequest("POST", "/sendmail", strings.NewReader(req2Body))
	req2.Header.Set("Content-Type", "application/json")
	resp2 := httptest.NewRecorder()
	router.ServeHTTP(resp2, req2)
	assert.Equal(t, http.StatusOK, resp2.Code)
	assert.Contains(t, resp2.Body.String(), `"code":1`)
	assert.Contains(t, resp2.Body.String(), `"msg":"ok"`)

	// Test case 3: email sending failed
	req3Body := `{"email": "test3@example.com"}`
	req3, _ := http.NewRequest("POST", "/sendmail", strings.NewReader(req3Body))
	req3.Header.Set("Content-Type", "application/json")
	resp3 := httptest.NewRecorder()
	router.ServeHTTP(resp3, req3)
	assert.Equal(t, http.StatusOK, resp3.Code)
	assert.Contains(t, resp3.Body.String(), `"code":1`)
	assert.Contains(t, resp3.Body.String(), `"msg":"ok"`)
}
