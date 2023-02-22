package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	tzh "github.com/go-playground/validator/v10/translations/zh"
)

type Validator interface {
	Validate(any) error
}

type validate struct {
	valid *validator.Validate
	tran  ut.Translator
}

// Validate 校验参数
func (vd *validate) Validate(v any) error {
	err := vd.valid.Struct(v)
	if ve, ok := err.(validator.ValidationErrors); ok {
		trans := ve.Translate(vd.tran)
		return &TranError{trans: trans, valid: ve}
	}
	return err
}

func (vd *validate) register(fn func(v *validator.Validate, t ut.Translator)) {
	fn(vd.valid, vd.tran)
}

func New() Validator {
	zht := zh.New()
	ent := en.New()
	uni := ut.New(ent, zht)

	tran, _ := uni.GetTranslator(zht.Locale())

	fn := func(field reflect.StructField) string {
		tag := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if tag == "" {
			tag = strings.SplitN(field.Tag.Get("query"), ",", 2)[0]
		}
		if tag == "" {
			tag = strings.SplitN(field.Tag.Get("form"), ",", 2)[0]
		}
		if tag == "" {
			tag = strings.SplitN(field.Tag.Get("yaml"), ",", 2)[0]
		}
		if tag == "-" {
			return ""
		}
		return tag
	}

	v := validator.New()
	v.RegisterTagNameFunc(fn)
	_ = tzh.RegisterDefaultTranslations(v, tran)

	vd := &validate{valid: v, tran: tran}
	// 自定义校验规则
	vd.register(substanceNameFunc)
	vd.register(usernameFunc)
	vd.register(passwordFunc)
	vd.register(hostnamePortFunc)
	vd.register(hostnameRFC1123Func)
	vd.register(hostnameFunc)
	vd.register(filenameFunc)
	vd.register(filepathFunc)
	vd.register(wsURLFunc)
	vd.register(httpURLFunc)
	vd.register(mobileFunc)
	vd.register(dongFunc)
	vd.register(tagFunc)
	vd.register(uniqueFunc)
	vd.register(requiredIfFunc)
	// vd.register(requiredWithoutFunc)
	vd.register(semverFunc)

	return vd
}
