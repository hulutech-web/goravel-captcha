package facades

import (
	"github.com/hulutech-web/goravel-captcha"
	"github.com/hulutech-web/goravel-captcha/contracts"
	"log"
)

func Captcha() contracts.Captcha {
	instance, err := captcha.App.Make(captcha.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Captcha)
}
