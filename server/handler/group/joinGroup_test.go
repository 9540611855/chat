package group

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

func TestJoinGroupHandler(t *testing.T) {
	init_.InitDB()
	// Create a new Gin router instance
	router := gin.Default()
	// Create a new group
	router.POST("/join-group", JoinGroup{}.JoinGroupHandler)
	// Join the group with valid data
	joinData := JoinGroup{
		Token:     "valid_jwt_token",
		UserId:    2,
		GroupId:   16823440491482086,
		TypeId:    1,
		GroupName: "Test Group",
	}
	joinJson, _ := json.Marshal(joinData)
	joinReq, _ := http.NewRequest("POST", "/join-group", bytes.NewBuffer(joinJson))
	joinReq.Header.Set("Content-Type", "application/json")
	joinResp := httptest.NewRecorder()
	router.ServeHTTP(joinResp, joinReq)

	assert.Equal(t, http.StatusOK, joinResp.Code)
	assert.Contains(t, joinResp.Body.String(), "请求已发送，等待管理员同意")
}
