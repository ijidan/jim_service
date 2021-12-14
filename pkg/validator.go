package pkg

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func MobileValidator(f validator.FieldLevel) bool {
	value := f.Field().String()
	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, value)
	if result {
		return true
	} else {
		return false
	}

}
