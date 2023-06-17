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

func TestFlashFriendHandle(t *testing.T) {
	init_.InitDB()
	// gin test router
	r := gin.Default()
	r.POST("/flash/message", FlashMessage{}.FlashFriendHandle)

	// test data
	flashMessage := FlashMessage{
		Token:  "valid_token",
		UserId: 4,
		Time:   1619600902,
	}

	// test request
	reqBody, _ := json.Marshal(flashMessage)
	req, err := http.NewRequest(http.MethodPost, "/flash/message", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// test response
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
