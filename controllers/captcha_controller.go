package controllers

import (
	"github.com/goravel/framework/contracts/http"
	Capt "github.com/hulutech-web/goravel-captcha/instance"
)

type CaptchaController struct {
}

func NewCaptchaController() *CaptchaController {
	return &CaptchaController{}
}

func (c *CaptchaController) GetCaptcha(ctx http.Context) http.Response {
	dots, b64, tb64, key, err := Capt.MakeCaptcha()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"errors": err,
		})
	}
	return ctx.Response().Success().Json(http.Json{
		"image":       b64,
		"thumb_image": tb64,
		"key":         key,
		"dots":        dots,
	})
}

func (c *CaptchaController) PostCaptcha(ctx http.Context) http.Response {
	type CaptchaReq struct {
		Dots string `json:"dots"`
		Key  string `json:"key"`
	}
	var req CaptchaReq
	if err := ctx.Request().Bind(&req); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "参数错误",
		})
	}

	checked := Capt.VerifyCaptcha(req.Key, req.Dots)
	if checked {
		return ctx.Response().Success().Json(http.Json{
			"message": "验证成功",
			"code":    200,
		})
	}
	return ctx.Response().Json(http.StatusInternalServerError, http.Json{
		"error": "验证失败",
	})
}
