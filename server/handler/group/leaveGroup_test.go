package group

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	init_ "pipiChat/init"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLeaveGroupHandler(t *testing.T) {
	init_.InitDB()
	// Create a new Gin router instance
	router := gin.Default()
	// Create a new group
	router.POST("/leave-group", LeaveGroup{}.LeaveGroupHandler)
	// Join the group with valid data
	// Leave the group with valid data
	leaveData := []byte(`{
		"token":   "valid_jwt_token",
		"userId":  1234,
		"groupId": 16823445656953373
	}`)
	leaveReq, _ := http.NewRequest("POST", "/leave-group", bytes.NewBuffer(leaveData))
	leaveReq.Header.Set("Content-Type", "application/json")
	leaveResp := httptest.NewRecorder()
	router.ServeHTTP(leaveResp, leaveReq)

	assert.Equal(t, http.StatusOK, leaveResp.Code)
	assert.Contains(t, leaveResp.Body.String(), "已退出群组")
}
