package captcha

import "github.com/mojocn/base64Captcha"

var store = base64Captcha.DefaultMemStore

// 生成验证码
func GenerateCaptcha() (id string, b64s string, answer string, err error) {
	/**
	NewDriverDigit(height,width,length,maxSkew,dotCount)
	验证码图片高度（单位：px）
	验证码图片宽度（单位：px）
	验证码字符个数（如：5 表示 5 位数字）
	最大倾斜角度（字符的扭曲程度，建议在 0~1 之间）
	背景干扰点数量（增加难度，防 OCR）
	*/
	driver := base64Captcha.NewDriverDigit(60, 180, 4, 0.7, 80)
	c := base64Captcha.NewCaptcha(driver, store)
	return c.Generate()
}

// 验证验证码
func VerifyCaptcha(id, answer string) bool {
	return store.Verify(id, answer, true)
}
