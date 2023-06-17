package friend

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pipiChat/configs"
	init_ "pipiChat/init"
	model "pipiChat/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddFriendHandle(t *testing.T) {
	init_.InitDB()

	// Call the EditPasswordHandle function with the HTTP request and response recorder
	r := gin.Default()
	r.POST("/addFriend", AddFriend{}.AddFriendHandle)
	user1 := model.User{UserName: "user1", Password: "1234", RealName: "下皮皮", Mail: "12345@qq.com"}
	user2 := model.User{UserName: "user2", Password: "1234", RealName: "皮皮", Mail: "12345@qq.com"}
	configs.DB.Create(&user1)
	configs.DB.Create(&user2)
	formData := []byte(`{
		"token":      "test_token",
		"userId":     4,
		"friendId":   27,
		"friendName": "t",
		"addType":    0
	}`)
	jsonValue, _ := json.Marshal(formData)
	t.Log(jsonValue)
	req, err := http.NewRequest("POST", "/addFriend", bytes.NewBuffer(formData))
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
