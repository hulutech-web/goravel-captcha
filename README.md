# goravel-captcha
基于[go-captcha](https://github.com/wenlng/go-captcha)的二次开发,

## 一、安装
```go
go get github.com/hulutech-web/goravel-captcha
```
## 二、配置
项目的`config/app.go`下
```go
"github.com/hulutech-web/goravel-captcha"
```

```go
"providers": []foundation.ServiceProvider{
    ...
    &captcha.ServiceProvider{},
}
```
## 三、功能概述
系统默认提供了2个验证码路由，一个是获取验证码，一个是校验验证码
扩展包captcha/routers/router.go
### 3.1 获取验证码
``
route.Get("api/captcha", captchaController.GetCaptcha)
``
### 3.2 校验验证码
``
route.Post("api/captcha", captchaController.PostCaptcha)
``
### 四、参数说明
### 4.1 获取验证码参数，返回类型`GetCaptchaRequest`
```go
type GetCaptchaRequest struct {
    Image      string               `json:"image"`       //base64图片数据
    ThumbImage string               `json:"thumb_image"` //缩略图数据，base64
    Key        string               `json:"key"`         //验证码唯一的key
    Dots       map[int]Capt.CharDot `json:"dots"`        //文字坐标点
}

type CharDot struct {
    // 顺序索引
    Index int
    // x,y位置
    Dx int
    Dy int
    // 字体大小
    Size int
    // 字体宽
    Width int
    // 字体高
    Height int
    // 字符文本
    Text string
    // 字体角度
    Angle int
    // 颜色
    Color string
    // 颜色2
    Color2 string
}
```

### 4.2 提交验证参数,返回类型`CaptchaReq`，该类型由前端配套组件[goravel-captcha-vue](https://github.com/hulutech-web/goravel-captcha-vue)提供，无需自行实现
```go
type CaptchaReq struct {
		Dots string `json:"dots"`
		Key  string `json:"key"`
	}
```
### 五、使用说明
captcha_controller已封装到扩展包中，代码如下：  
```go
package controllers

import (
	"github.com/goravel/framework/contracts/http"
	Capt "github.com/hulutech-web/goravel-captcha"
)

type CaptchaController struct {
}

func NewCaptchaController() *CaptchaController {
	return &CaptchaController{}
}

func (c *CaptchaController) GetCaptcha(ctx http.Context) http.Response {
	type GetCaptchaRequest struct {
		Image      string               `json:"image"`       //base64图片数据
		ThumbImage string               `json:"thumb_image"` //缩略图数据，base64
		Key        string               `json:"key"`         //验证码唯一的key
		Dots       map[int]Capt.CharDot `json:"dots"`        //文字坐标点
	}
	dots, b64, tb64, key, err := Capt.MakeCaptcha()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"errors": err,
		})
	}
	var req GetCaptchaRequest
	req.Dots = dots.(map[int]Capt.CharDot)
	req.Image = b64
	req.ThumbImage = tb64
	req.Key = key

	return ctx.Response().Success().Json(req)
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
```
### 六、预览效果，前端由[goravel-captcha-vue](https://github.com/hulutech-web/goravel-captcha)提供，无需自行实现
![image](https://github.com/hulutech-web/goravel-captcha/blob/master/images/default.png?raw=true)
![image](https://github.com/hulutech-web/goravel-captcha/blob/master/images/success.png?raw=true)
![image](https://github.com/hulutech-web/goravel-captcha/blob/master/images/validating.png?raw=true)
![image](https://github.com/hulutech-web/goravel-captcha/blob/master/images/validated.png?raw=true)

