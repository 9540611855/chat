package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// configJsonBody json request body.
type configJsonBody struct {
	Id          string
	CaptchaType string
	VerifyValue string
}
type driverJson struct {
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

// 	{
// 		ShowLineOptions: [],
// 		CaptchaType: "string",
// 		Id: '',
// 		VerifyValue: '',
// 		DriverAudio: {
// 			Length: 6,
// 			Language: 'zh'
// 		},
// 		DriverString: {
// 			Height: 60,
// 			Width: 240,
// 			ShowLineOptions: 0,
// 			NoiseCount: 0,
// 			Source: "1234567890qwertyuioplkjhgfdsazxcvbnm",
// 			Length: 6,
// 			Fonts: ["wqy-microhei.ttc"],
// 			BgColor: {R: 0, G: 0, B: 0, A: 0},
// 		},
// 		DriverMath: {
// 			Height: 60,
// 			Width: 240,
// 			ShowLineOptions: 0,
// 			NoiseCount: 0,
// 			Length: 6,
// 			Fonts: ["wqy-microhei.ttc"],
// 			BgColor: {R: 0, G: 0, B: 0, A: 0},
// 		},
// 		DriverChinese: {
// 			Height: 60,
// 			Width: 320,
// 			ShowLineOptions: 0,
// 			NoiseCount: 0,
// 			Source: "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,,不想要,的值",
// 			Length: 2,
// 			Fonts: ["wqy-microhei.ttc"],
// 			BgColor: {R: 125, G: 125, B: 0, A: 118},
// 		},
// 		DriverDigit: {
// 			Height: 80,
// 			Width: 240,
// 			Length: 5,
// 			MaxSkew: 0.7,
// 			DotCount: 80
// 		}
// 	},
// 	blob: "",
// 	loading: false
// }

var store = base64Captcha.DefaultMemStore
var tokenStore map[string]struct{}

// base64Captcha create http handler
func GenerateCaptchaHandler(c *gin.Context) {
	//parse request parameters
	decoder := json.NewDecoder(c.Request.Body)
	var param configJsonBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer c.Request.Body.Close()
	var driver base64Captcha.Driver
	driverJsons := driverJson{
		DriverAudio: &base64Captcha.DriverAudio{},
		DriverString: &base64Captcha.DriverString{
			Length:          4,
			Height:          60,
			Width:           240,
			ShowLineOptions: 2,
			NoiseCount:      0,
			Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		},
		DriverChinese: &base64Captcha.DriverChinese{},
		DriverMath:    &base64Captcha.DriverMath{},
		DriverDigit:   &base64Captcha.DriverDigit{},
	}

	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = driverJsons.DriverAudio
	case "string":
		driver = driverJsons.DriverString.ConvertFonts()
	case "math":
		driver = driverJsons.DriverMath.ConvertFonts()
	case "chinese":
		driver = driverJsons.DriverChinese.ConvertFonts()
	default:
		driver = driverJsons.DriverDigit
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	c.JSON(http.StatusOK, body)
}

// base64Captcha verify http handler
func CaptchaVerifyHandle(c *gin.Context) {

	//parse request json body
	decoder := json.NewDecoder(c.Request.Body)
	var param configJsonBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer c.Request.Body.Close()
	//verify the captcha
	body := map[string]interface{}{"code": 0, "msg": "failed"}
	if store.Verify(param.Id, param.VerifyValue, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}

	//set json response
	c.JSON(http.StatusOK, body)
}
