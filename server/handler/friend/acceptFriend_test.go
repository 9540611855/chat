package friend

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	init_ "pipiChat/init"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAcceptFriendHandle(t *testing.T) {
	init_.InitDB()

	// Call the EditPasswordHandle function with the HTTP request and response recorder
	r := gin.Default()
	r.POST("/accept_Friend_test", AcceptFriend{}.AcceptFriendHandle)
	formData := []byte(`{
		"token":      "test_token",
		"userId":     4,
		"friendId":   27
	}`)
	//jsonValue, _ := json.Marshal(formData)
	//t.Log(jsonValue)
	req, err := http.NewRequest("POST", "/accept_Friend_test", bytes.NewBuffer(formData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//判断返回结果
	assert.Equal(t, http.StatusOK, w.Code)
	var response gin.H
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "success")
	assert.NotContains(t, response, "error")
}
