package initialize

import (
	"chuckie_gin/global"
	"chuckie_gin/utils"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func InitTrans (locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 获取tag form:"password" json:"password" binding:"required,min=3,max=20"
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()

		uni := ut.New(enT, zhT, enT)

		global.Trans, ok = uni.GetTranslator(locale)

		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		}
		// 注册手机号校验
		RegisterValidatorFunc(v, "mobile", "手机号非法", utils.ValidateMobile)
		return
	}
	return
}

type Func func(fl validator.FieldLevel) bool

func RegisterValidatorFunc(v *validator.Validate, tag string, msgStr string, fn Func) {
	_ = v.RegisterValidation(tag, validator.Func(fn))

	v.RegisterTranslation(tag, global.Trans, func(ut ut.Translator) error {
		return ut.Add(tag, "{0}"+msgStr, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return  t
	})
}
