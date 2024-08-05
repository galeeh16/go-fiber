package helper

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type ErrorResponse struct {
	FailedField string `json:"field" label:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
	Message     string `json:"message"`
}

var validate = validator.New()

func ValidateStruct(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(data)

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Message = "This is the message"
			errors = append(errors, &element)
		}
	}

	fmt.Println(errors)

	// errs := translateError(errors, trans)
	return errors
	// return errs
}

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
