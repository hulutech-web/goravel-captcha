package captcha

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/hulutech-web/goravel-captcha/routers"
)

const Binding = "captcha"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return nil, nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("./packages/captcha", map[string]string{
		"config/captcha.go": app.ConfigPath("captcha.go"),
	})
	//初始化路由
	routers.InitCaptcha(app)
}
