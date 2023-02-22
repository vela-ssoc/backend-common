package validate

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// TranError 翻译后的错误信息
type TranError struct {
	trans validator.ValidationErrorsTranslations
	valid validator.ValidationErrors
}

func (e *TranError) Error() string {
	causes := make([]string, 0, len(e.valid))
	for _, val := range e.valid {
		ns := val.Namespace()
		cause := e.trans[ns]
		causes = append(causes, cause)
	}
	return strings.Join(causes, ", ")
}
