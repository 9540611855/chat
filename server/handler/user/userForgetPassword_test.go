package user

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"pipiChat/handler/auth"
	init_ "pipiChat/init"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestForgePasswordHandle(t *testing.T) {
	// create a new gin router
	r := gin.Default()
	init_.InitDB()
	// define the route
	r.POST("/forge_password", ForgePasswordHandle)

	// create a test request
	jsonStr := []byte(`{
        "email": "954061185@qq.com",
        "verification_code": "123456",
        "password": "newpassword",
        "user_id": 1
    }`)
	var userReg auth.UserEmailRegistration
	userReg.Email = "954061185@qq.com"
	userReg.VerificationCode = "123456"
	auth.MailstoreArray = append(auth.MailstoreArray, userReg)

	req, err := http.NewRequest("POST", "/forge_password", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create a test response recorder
	rr := httptest.NewRecorder()

	// perform the request
	r.ServeHTTP(rr, req)

	// check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check the response body
	expected := `{"success":"修改密码成功"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
