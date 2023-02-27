package netutil

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/vela-ssoc/backend-common/pubody"
	"github.com/xgfone/ship/v5"
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
	code := http.StatusNotFound
	ret := &pubody.BizError{Code: code, Node: nh.node, Cause: "资源不存在"}
	return c.JSON(code, ret)
}

func (nh *nodeHandle) Error(c *ship.Context, e error) {
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
	ret := &pubody.BizError{Code: code, Node: nh.node, Cause: cause}
	_ = c.JSON(code, ret)
}
