package main

import (
	"fmt"

	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ru_translations "github.com/go-playground/validator/v10/translations/ru"
)
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)
func main()  {
	en := english.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("ru")

	validate := validator.New()
	ru_translations.RegisterDefaultTranslations(validate, trans)

	myVar := "12"
	err := validate.Var(myVar,"required,min=3")
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs)
		for _, e := range errs {
			fmt.Printf("%s\n",e.Translate(trans))
		}
	}

}

