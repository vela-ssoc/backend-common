package problem

import (
	"net/http"
	"time"

	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

// Handler ship 框架的一些错误处理
type Handler interface {
	// NotFound 路由不存在的处理方法
	NotFound(*ship.Context) error

	// HandleError 错误统一处理方法
	HandleError(*ship.Context, error)
}

func NewHandle(name string) Handler {
	if name == "" {
		name = "about:blank"
	}
	return &handle{name: name}
}

type handle struct {
	name string
}

func (h *handle) NotFound(*ship.Context) error {
	return ship.ErrNotFound.Newf("资源不存在")
}

func (h *handle) HandleError(c *ship.Context, e error) {
	ret := &Problem{
		Type:     h.name, // 一般为 URL 全路径，此处设置成节点名字，也便于排查问题
		Title:    "请求错误",
		Status:   http.StatusBadRequest,
		Detail:   e.Error(),
		Instance: c.RequestURI(),
	}

	switch err := e.(type) {
	case ship.HTTPServerError:
		ret.Status = err.Code
	case *time.ParseError:
		ret.Title = "参数错误"
		ret.Detail = "时间格式错误（正确格式：" + err.Layout + "）"
	default:
		switch {
		case err == gorm.ErrRecordNotFound:
			ret.Detail = "数据不存在"
		}
	}

	// 按照 RFC7807 规范 Content-Type 设置为 application/problem+json 更为妥当，
	// 此处偷个懒使用默认 application/json。
	_ = c.JSON(ret.Status, ret)
}
