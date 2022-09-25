package utility

import (
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func GetNewCaptcha() (string, string, error) {
	var driver base64Captcha.Driver
	driver = base64Captcha.DefaultDriverDigit
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	return id, b64s, err
}

func VerifyCaptcha(id string, code string) bool {
	return store.Verify(id, code, true)
}
