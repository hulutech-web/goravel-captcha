package routers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/hulutech-web/goravel-captcha/controllers"
)

func InitCaptcha(app foundation.Application) {
	route := app.MakeRoute()
	captchaController := controllers.NewCaptchaController()
	route.Get("api/captcha", captchaController.GetCaptcha)
	route.Post("api/captcha", captchaController.PostCaptcha)
}
