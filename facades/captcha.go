package facades

import (
	"goravel/packages/captcha"
	"goravel/packages/captcha/contracts"
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
