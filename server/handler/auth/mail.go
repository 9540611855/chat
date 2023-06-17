package auth

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"pipiChat/configs"
	model "pipiChat/models"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type UserEmailRegistration struct {
	Email            string `json:"email"`
	VerificationCode string `json:"verification_code"`
}

var MailstoreArray []UserEmailRegistration

// var MailstoreArray = make(MailstoreArray, 40)
func registerEmail(param UserEmailRegistration) {
	for i, item := range MailstoreArray {
		if item.Email == param.Email {
			// 如果找到相同的 Email，则直接替换 VerificationCode
			MailstoreArray[i].VerificationCode = param.VerificationCode
			return
		}
	}
	// 如果没有找到相同的 Email，则将新的结构体添加到数组中
	MailstoreArray = append(MailstoreArray, param)
}

func sendMail(param UserEmailRegistration) bool {
	done := make(chan error)
	auth := smtp.PlainAuth("", "", "", "smtp.qq.com")
	// Compose the message.
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(param.Email) {
		return false
	}
	to := []string{param.Email}
	contentType := "Content-Type: text/html; charset=UTF-8"
	code := rand.Intn(1000000)
	param.VerificationCode = fmt.Sprintf("%06d", code)
	codeStr := fmt.Sprintf("欢迎注册pipiChat,验证码是:%06d\n\r", code)
	msgs := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s",
		to, "pipichat", "", "Subject", contentType, codeStr)
	msg := []byte(msgs)
	err := smtp.SendMail("smtp.qq.com:587", auth, "", to, msg)
	// Send the message.

	go func() {
		err := smtp.SendMail("smtp.qq.com:587", auth, "", to, msg)
		done <- err
	}()
	select {
	case err := <-done:
		if err != nil {
			return false
		}
	case <-time.After(2 * time.Minute):
		return false
	}

	if err != nil {
		return false
	}
	if len(MailstoreArray) > 38 {
		return false
	}
	registerEmail(param)
	return true
}

func ValueCode(param UserEmailRegistration) bool {
	for i := 0; i < len(MailstoreArray); i++ {
		if MailstoreArray[i].Email == param.Email {
			valueCode := MailstoreArray[i].VerificationCode

			MailstoreArray = append(MailstoreArray[:i], MailstoreArray[i+1:]...)
			if valueCode != param.VerificationCode {
				return false
			} else {

				return true
			}
		}
	}
	return false
}

func ValueEmail(param UserEmailRegistration) bool {
	var selectUser model.User
	result := configs.DB.Find(&selectUser, "mail = ?", param.Email)
	if result.RowsAffected != 0 {
		return true
	}
	return false
}

func ValueCodeHandle(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)
	var param UserEmailRegistration
	err := decoder.Decode(&param)
	var body = map[string]interface{}{"msg": "ok"}
	if err != nil {
		body = map[string]interface{}{"msg": "error"}
		c.JSON(http.StatusBadRequest, body)
		return
	}
	if ValueEmail(param) {
		body = gin.H{"error": "这个邮箱没有注册"}
		c.JSON(http.StatusBadRequest, body)
		return
	}

	for i := 0; i < len(MailstoreArray); i++ {
		if MailstoreArray[i].Email == param.Email {
			valueCode := MailstoreArray[i].VerificationCode

			MailstoreArray = append(MailstoreArray[:i], MailstoreArray[i+1:]...)
			if valueCode != param.VerificationCode {
				body = gin.H{"error": "验证码错误"}
				c.JSON(http.StatusBadRequest, body)
				return
			} else {
				//todo 插入regmali表里面数据
				c.JSON(http.StatusOK, body)
				return
			}
		}
	}
	body = gin.H{"error": "这个邮箱没有注册"}
	c.JSON(http.StatusBadRequest, body)
	return

}

func SendMailHandle(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)
	var param UserEmailRegistration
	err := decoder.Decode(&param)
	var body = gin.H{"success": "发送邮箱成功,请注意接受"}
	if err != nil {
		body = gin.H{"error": "系统错误"}
		c.JSON(http.StatusBadRequest, body)
		return
	}
	res := sendMail(param)
	if !res {
		body = gin.H{"error": "系统错误"}
		c.JSON(http.StatusBadRequest, body)
		return
	}
	//set json response
	c.JSON(http.StatusOK, body)
}
