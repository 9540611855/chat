package user

import (
	"net/http"
	"net/http/httptest"
	init_ "pipiChat/init"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEditPasswordHandle(t *testing.T) {
	init_.InitDB()
	// Create a new HTTP request
	reqBody := `{"jwt_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2ODA3NzM1OTMsInVzZXIiOjF9.LD5LodtmhxPatSft7pUfmtN3Y-PVXJRIrolCnHxsbTo", "password": "new_password"}`
	req := httptest.NewRequest("POST", "/edit-password", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Call the EditPasswordHandle function with the HTTP request and response recorder
	router := gin.Default()
	router.POST("/edit-password", EditPasswordHandle)
	router.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	expectedBody := `{"success":"修改密码成功"}`
	if w.Body.String() != expectedBody {
		t.Errorf("expected response body %s but got %s", expectedBody, w.Body.String())
	}
}
