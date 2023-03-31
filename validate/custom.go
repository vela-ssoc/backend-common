package validate

import (
	"net/url"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	unt "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// uniqueFunc unique 翻译
func uniqueFunc(v *validator.Validate, t unt.Translator) {
	const tag = "unique"
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}必须唯一", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// requiredIfFunc required_if 翻译
func requiredIfFunc(v *validator.Validate, t unt.Translator) {
	const tag = "required_if"
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}为{1}时{2}必须填写", true)
	}, func(ut unt.Translator, fe validator.FieldError) (str string) {
		if sn := strings.SplitN(fe.Param(), " ", 2); len(sn) == 2 {
			str, _ = ut.T(tag, strings.ToLower(sn[0]), sn[1], fe.Field())
		} else {
			str, _ = ut.T(tag, "<undefined>", "<undefined>", fe.Field())
		}
		return
	})
}

// requiredWithoutFunc required_without 翻译
//func requiredWithoutFunc(v *validator.Validate, t unt.Translator) {
//	const tag = "required_without"
//	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
//		return ut.Add(tag, "{0}不存在时{1}必须填写", true)
//	}, func(ut unt.Translator, fe validator.FieldError) string {
//		target := fe.Param()
//		str, _ := ut.T(tag, strings.ToLower(target), fe.Field())
//		return str
//	})
//}

// semverFunc 语义化版本号：https://semver.org/lang/zh-CN/
func semverFunc(v *validator.Validate, t unt.Translator) {
	const tagName = "semver"
	_ = v.RegisterTranslation(tagName, t, func(ut unt.Translator) error {
		return ut.Add(tagName, "{0}必须是语义化版本号", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tagName, fe.Field())
		return s
	})
}

// hostnamePortFunc hostname_port 翻译
func hostnamePortFunc(v *validator.Validate, t unt.Translator) {
	const tag = "hostname_port"
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}必须包含主机名与端口号", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// hostnameRFC1123Func https://www.rfc-editor.org/rfc/rfc1123.html
func hostnameRFC1123Func(v *validator.Validate, t unt.Translator) {
	const tag = "hostname_rfc1123"
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}必须是一个域名", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// mobileFunc 中国大陆手机号校验器
func mobileFunc(v *validator.Validate, t unt.Translator) {
	// https://www.jianshu.com/p/1e8eab706a63
	reg := regexp.MustCompile("^1(3\\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\\d|9[0-35-9])\\d{8}$")
	tag := "mobile"

	_ = v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		field := fl.Field()
		if field.Kind() == reflect.String {
			return reg.MatchString(field.String())
		}
		return false
	})
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}必须是中国大陆手机号", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// httpURLFunc 必须是 http/https 协议的 URL
func httpURLFunc(v *validator.Validate, t unt.Translator) {
	const tag = "http_url"
	_ = v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			pu, _ := url.Parse(field.String())
			return pu != nil && (strings.EqualFold("http", pu.Scheme) || strings.EqualFold("https", pu.Scheme))
		}
		return false
	})
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}必须是http(s)协议的url", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// wsURLFunc 必须是 ws/wss 协议的 URL
func wsURLFunc(v *validator.Validate, t unt.Translator) {
	const tag = "ws_url"
	_ = v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			pu, _ := url.Parse(field.String())
			return pu != nil && (strings.EqualFold("ws", pu.Scheme) || strings.EqualFold("wss", pu.Scheme))
		}
		return false
	})
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}必须是ws(s)协议的url", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// usernameFunc 用户名校验
func usernameFunc(v *validator.Validate, t unt.Translator) {
	const tag = "username"
	reg := regexp.MustCompile("^[a-zA-Z0-9._-]{2,20}$")
	_ = v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			return reg.MatchString(field.String())
		}
		return false
	})
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}必须是2-20位的数字、字母、减号、下划线组合", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// hostnameFunc 主机名校验器翻译
func hostnameFunc(v *validator.Validate, t unt.Translator) {
	const tag = "hostname"
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}必须是一个主机名", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// passwordFunc 密码校验
func passwordFunc(v *validator.Validate, t unt.Translator) {
	const tag = "password"
	_ = v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			passwd := field.String()
			length := len(passwd)
			if length < 8 || length > 32 { // 密码8-32位
				return false
			}
			// 请参考 ASCII 码表，一目了然
			isLower := func(u rune) bool { return u >= 'a' && u <= 'z' }
			isUpper := func(u rune) bool { return u >= 'A' && u <= 'Z' }
			isNumber := func(u rune) bool { return u >= '0' && u <= '9' }
			isOther := func(u rune) bool {
				return (u >= '!' && u <= '/') ||
					(u >= ':' && u <= '@') ||
					(u >= '[' && u <= '`') ||
					(u >= '{' && u <= '~')
			}

			// 必须包含小写字母 大写字母 数字 特殊字符，且不能有其他字符
			var lower, upper, number, other bool
			for _, u := range passwd {
				if isLower(u) {
					lower = true
				} else if isUpper(u) {
					upper = true
				} else if isNumber(u) {
					number = true
				} else if isOther(u) {
					other = true
				} else {
					return false // 不符合密码要求的四种字符
				}
			}

			return lower && upper && number && other
		}
		return false
	})
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}不符合密码强度要求", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

func substanceNameFunc(v *validator.Validate, t unt.Translator) {
	const tag = "substance_name"
	reg := regexp.MustCompile("^[a-zA-Z0-9-_@#\u4e00-\u9fa5]{1,20}$")

	_ = v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			return reg.MatchString(field.String())
		}
		return false
	})
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}不符合规范", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// filenameFunc 文件名校验器
func filenameFunc(v *validator.Validate, t unt.Translator) {
	const tag = "filename"
	// reg := regexp.MustCompile("^[a-zA-Z0-9\u4e00-\u9fa5-_.]{1,40}$")
	reg := regexp.MustCompile("^[a-zA-Z0-9._-]{1,50}$")
	_ = v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		f := fl.Field()
		if f.Kind() != reflect.String {
			return false
		}
		// 文件名不能以 . 开头
		str := f.String()
		if str == "" || str[0] == '.' {
			return false
		}
		// 要符合正则表达式
		if ok := reg.MatchString(str); !ok {
			return false
		}
		// 文件名不能是系统特殊含义字符
		ext := filepath.Ext(str) // 获取文件后缀
		if ext == ".zip" {       // zip 文件解压后的名字也不能是系统特殊含义字符
			str = str[:len(str)-4]
		}

		return !isReservedName(str)
	})
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}不是一个合法的文件名", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// dongFunc 咚咚校验器
func dongFunc(v *validator.Validate, t unt.Translator) {
	const tag = "dong"

	// 集团咚咚号为6位纯数字
	// 证券咚咚号为5位纯属字
	reg := regexp.MustCompile("^[0-9]{5,6}$")

	_ = v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			return reg.MatchString(field.String())
		}
		return false
	})
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}必须是一个有效的咚咚号", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// tagFunc 节点标签校验
func tagFunc(v *validator.Validate, t unt.Translator) {
	const tag = "tag"
	reg := regexp.MustCompile("^[\u4e00-\u9fa5a-zA-Z0-9@._-]{2,15}$")

	_ = v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			return reg.MatchString(field.String())
		}
		return false
	})
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}标签命名不规范", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// filepathFunc 文件路径校验
func filepathFunc(v *validator.Validate, t unt.Translator) {
	const tag = "filepath"
	reg := regexp.MustCompile("^[a-zA-Z0-9./_\u4e00-\u9fa5\\\\]{3,50}$") // 一-龥
	_ = v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			str := field.String()
			if !reg.MatchString(str) {
				return false
			}
			return !containsReservedName(str)
		}
		return false
	})
	_ = v.RegisterTranslation(tag, t, func(ut unt.Translator) error {
		return ut.Add(tag, "{0}必须是一个合法的文件路径", true)
	}, func(ut unt.Translator, fe validator.FieldError) string {
		s, _ := ut.T(tag, fe.Field())
		return s
	})
}

// containsReservedName 判断路径中是否含有系统保留字
func containsReservedName(path string) bool {
	clean := filepath.Clean(path)
	format := strings.ReplaceAll(clean, "\\", "/")
	splits := strings.Split(format, "/")
	for _, ele := range splits {
		if isReservedName(ele) {
			return true
		}
	}
	return false
}

// reservedNames lists reserved Windows names. Search for PRN in
// https://docs.microsoft.com/en-us/windows/desktop/fileio/naming-a-file
// for details.
var reservedNames = []string{
	"CON", "PRN", "AUX", "NUL",
	"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
	"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
}

// isReservedName returns true, if path is Windows reserved name.
// See reservedNames for the full list.
func isReservedName(path string) bool {
	if len(path) == 0 {
		return false
	}
	for _, reserved := range reservedNames {
		if strings.EqualFold(path, reserved) {
			return true
		}
	}
	return false
}
