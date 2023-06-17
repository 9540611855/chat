// Test generateCaptchaHandler
package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGenerateCaptchaHandler(t *testing.T) {
	r := gin.Default()
	r.POST("/captcha/generate", GenerateCaptchaHandler)

	reqBody := configJsonBody{
		CaptchaType: "string",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)
	//t.Log(reqBodyBytes)
	req, err := http.NewRequest("POST", "/captcha/generate", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response map[string]interface{}
	json.Unmarshal([]byte(rr.Body.String()), &response)
	t.Log(rr)

	if (int64(response["code"].(float64))) != 1 {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
	t.Log("\n\n\n\n")
	res := store.Get(response["captchaId"].(string), false)
	t.Log(res)
}

// Test captchaVerifyHandle
func TestCaptchaVerifyHandle(t *testing.T) {
	r := gin.Default()
	r.POST("/captcha/verify", CaptchaVerifyHandle)

	reqBody := configJsonBody{
		Id:          "jcz2C9Ue42JUmnUYghyy",
		VerifyValue: "yzke",
	}

	reqBodyBytes, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/captcha/verify", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response map[string]interface{}
	json.Unmarshal([]byte(rr.Body.String()), &response)

	if (int64(response["code"].(float64))) != 1 {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}
