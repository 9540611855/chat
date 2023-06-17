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

func TestCreateGroupHandler(t *testing.T) {
	init_.InitDB()
	// Create a new Gin router instance
	router := gin.Default()

	// Add the CreateGroupHandler function as a route handler for the POST /create-group endpoint
	router.POST("/create-group", CreateGroupHandler)

	// Create a new request body
	body := &CreateGroup{
		Token:     "test_token",
		UserId:    1234,
		GroupName: "Test Group",
	}
	jsonBody, _ := json.Marshal(body)

	// Create a new HTTP POST request with the request body
	req, err := http.NewRequest("POST", "/create-group", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set the content type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a new recorder to capture the response
	respRecorder := httptest.NewRecorder()

	// Call the router's ServeHTTP method with the request and response recorder
	router.ServeHTTP(respRecorder, req)

	// Check that the response status code is OK
	assert.Equal(t, http.StatusOK, respRecorder.Code)

}
