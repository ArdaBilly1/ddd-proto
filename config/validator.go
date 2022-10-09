package config

import (
	"errors"
	"reflect"

	"github.com/go-playground/locales/id"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

type Validator struct {
	Validator *validator.Validate
}

var Translate translator.Translator

func (v *Validator) Validate(i interface{}) error {
	id := id.New()
	uni := translator.New(id, id)

	// translate to bahasa
	var ok bool
	Translate, ok = uni.GetTranslator("id")
	if !ok {
		return errors.New("can't find translator")
	}

	err := id_translations.RegisterDefaultTranslations(v.Validator, Translate)
	if err != nil {
		return errors.New("can't register translator")
	}

	v.Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})

	return v.Validator.Struct(i)
}
