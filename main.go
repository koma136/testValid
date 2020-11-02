package main

import (
	"fmt"
	ru "github.com/go-playground/locales/ru_RU"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func main()  {
	var uni      *ut.UniversalTranslator
	ru := ru.New()
	uni = ut.New(ru, ru)
	trans, _ := uni.GetTranslator("ru")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Var("12","required,min=3")
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Errorf("value: %s", err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Errorf("%s",err.Translate(trans))
		}
	}

}

