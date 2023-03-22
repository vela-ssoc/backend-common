package problem

import (
	"encoding/base64"
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/vela-ssoc/backend-common/validate"
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
	pd := &Detail{
		Type:     h.name, // 一般为 URL 全路径，此处设置成节点名字，也便于排查问题
		Title:    "请求错误",
		Status:   http.StatusBadRequest,
		Detail:   e.Error(),
		Instance: c.RequestURI(),
	}

	switch err := e.(type) {
	case ship.HTTPServerError:
		pd.Status = err.Code
	case *validate.TranError:
		pd.Title = "参数校验错误"
	case *time.ParseError:
		pd.Title = "参数格式错误"
		pd.Detail = "时间格式错误，正确格式：" + err.Layout
	case *net.ParseError:
		pd.Title = "参数格式错误"
		pd.Detail = err.Text + " 不是有效的 " + err.Type
	case base64.CorruptInputError:
		pd.Title = "参数格式错误"
		pd.Detail = "base64 编码错误：" + err.Error()
	case *json.SyntaxError:
		pd.Title = "报文格式错误"
		pd.Detail = "请求报错必须是 JSON 格式"
	case *json.UnmarshalTypeError:
		pd.Title = "数据类型错误"
		pd.Detail = err.Field + " 收到无效的数据类型"
	case *strconv.NumError:
		pd.Title = "数据类型错误"
		var msg string
		if sn := strings.SplitN(err.Func, "Parse", 2); len(sn) == 2 {
			msg = err.Num + " 不是 " + strings.ToLower(sn[1]) + " 类型"
		} else {
			msg = "类型错误：" + err.Num
		}
		pd.Detail = msg
	case *mysql.MySQLError:
		switch err.Number {
		case 1062:
			pd.Detail = "数据已存在"
		default:
			c.Errorf("SQL 执行错误：%v", e)
			pd.Status = http.StatusInternalServerError
			pd.Detail = "内部错误"
		}
	default:
		switch {
		case err == gorm.ErrRecordNotFound:
			pd.Detail = "数据不存在"
		}
	}

	_ = pd.JSON(c.ResponseWriter())
}
