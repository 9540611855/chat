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

func TestBlockFriendHandle(t *testing.T) {
	init_.InitDB()
	// Create request body
	blockFriend := BlockFriend{
		Token:    "fake jwt token",
		UserId:   4,
		FriendId: 27,
	}
	reqBody, _ := json.Marshal(blockFriend)

	// Create HTTP test server and client
	router := gin.Default()
	router.POST("/block-friend", blockFriend.BlockFriendHandle)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/block-friend", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Assert response status code and body
	assert.Equal(t, http.StatusOK, w.Code)
	var resBody gin.H
	err := json.Unmarshal(w.Body.Bytes(), &resBody)
	assert.NoError(t, err)
	assert.Equal(t, gin.H{"success": "拉黑好友成功"}, resBody)
}
