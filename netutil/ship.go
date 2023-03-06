package netutil

import (
	"net/http"
	"time"

	"github.com/vela-ssoc/backend-common/pubody"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func Notfound(node string) ship.Handler {
	nh := &nodeHandle{node: node}
	return nh.Notfound
}

func ErrorFunc(node string) func(*ship.Context, error) {
	nh := &nodeHandle{node: node}
	return nh.Error
}

type nodeHandle struct {
	node string
}

func (nh *nodeHandle) Notfound(c *ship.Context) error {
	r := c.Request()
	method, dst := r.Method, r.URL
	return ship.ErrNotFound.Newf("请求资源不存在：%s %s", method, dst)
}

func (nh *nodeHandle) Error(c *ship.Context, e error) {
	code, cause := http.StatusBadRequest, e.Error()

	switch err := e.(type) {
	case ship.HTTPServerError:
		code = err.Code
	case *time.ParseError:
		cause = "时间格式错误（正确格式：" + err.Layout + "）"
	default:
		switch {
		case err == gorm.ErrRecordNotFound:
			cause = "数据不存在"
		}
	}
	ret := &pubody.BizError{Code: code, Node: nh.node, Cause: cause}
	_ = c.JSON(code, ret)
}
