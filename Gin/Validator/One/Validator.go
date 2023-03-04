package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
)

var Trans ut.Translator

func Validator() gin.HandlerFunc {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhHans := zh_Hans.New()
		uni := ut.New(zhHans)

		Trans, _ = uni.GetTranslator("")
		zh.RegisterDefaultTranslations(v, Trans)

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("label")
		})

		v.RegisterTranslation("required", Trans, func(ut ut.Translator) error {
			return ut.Add("required", "请输入{0}", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T(fe.Tag(), fe.Field())
			if err != nil {
				return fe.(error).Error()
			}
			return t
		})
	}
	return func(c *gin.Context) {

	}
}

func ValidatorErrors(c *gin.Context, err error) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		errs := gin.H{}
		for _, e := range errors {
			errs[e.StructField()] = e.Translate(Trans)
		}
		c.JSON(http.StatusUnprocessableEntity, errs)
	} else {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
}
