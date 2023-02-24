package pubrr

import (
	"net/http"

	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func NotFound(node string) ship.Handler {
	return func(c *ship.Context) error {
		code := http.StatusNotFound
		ret := &ErrorResult{Code: code, Node: node, Cause: "404 not found"}
		c.SetContentType("application/json; charset=utf-8")
		return c.JSON(code, ret)
	}
}

func ErrorHandle(node string) func(*ship.Context, error) {
	return func(c *ship.Context, e error) {
		code, cause := http.StatusBadRequest, e.Error()

		switch err := e.(type) {
		case ship.HTTPServerError:
			code = err.Code
		default:
			switch {
			case err == gorm.ErrRecordNotFound:
				cause = "数据不存在"
			}
		}

		ret := &ErrorResult{Code: code, Node: node, Cause: cause}
		_ = c.JSON(code, ret)
	}
}
