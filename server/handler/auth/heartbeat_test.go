package auth

import (
	"net/http"
	"net/http/httptest"
	init_ "pipiChat/init"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHeartbeatHandler(t *testing.T) {
	// Create a new Gin router.
	router := gin.New()
	init_.InitDB()
	// Add the heartbeatHandler route.
	router.POST("/heartbeat", HeartbeatHandler)

	// Set the request body.
	jsonStr := `{"userid": 1, "token": "error", "jwttoken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2ODAwMDQ3ODksInVzZXIiOjF9.Ot52X47mDneX93Pj5Dc_7IP7VaYfY51mILUN7M-ob54"}`

	req1, _ := http.NewRequest("POST", "/heartbeat", strings.NewReader(jsonStr))
	req1.Header.Set("Content-Type", "application/json")
	resp1 := httptest.NewRecorder()
	router.ServeHTTP(resp1, req1)
	// Perform the request.
	// Check the response status code.
	if status := resp1.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body.
	expected := `{"success":"下发token成功"}`
	if resp1.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp1.Body.String(), expected)
	}
}

func TestRegenerateJwtToken(t *testing.T) {
	init_.InitDB()
	res := regenerateJwtToken(1)
	if !res {
		t.Fatalf("failed to generate old token: %v", res)
	}
}
