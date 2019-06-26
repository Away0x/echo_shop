package captcha

import (
	"echo_shop/config"
	captchaUtil "echo_shop/pkg/captcha"
)

// CreateCaptcha 生成验证码
func CreateCaptcha() captchaUtil.Captcha {
	return captchaUtil.New(func(id string) string {
		return config.Application.Reverse("captcha", id) // 生成验证码 url
	})
}

// Verify 验证
func Verify(captchaID, captchaVal string) bool {
	return captchaUtil.Verify(captchaID, captchaVal)
}
