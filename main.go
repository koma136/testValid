package main

import (
	"fmt"
	ru "github.com/go-playground/locales/ru_RU"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)
func main()  {
	ru := ru.New()
	uni = ut.New(ru, ru)
	trans, _ := uni.GetTranslator("ru")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Var("12","required,min=3")
	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			fmt.Errorf(e.Translate(trans))
		}
	}

}

