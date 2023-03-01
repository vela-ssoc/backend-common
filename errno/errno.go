package errno

import "github.com/xgfone/ship/v5"

func NodeNotfound(arg string) ship.HTTPServerError {
	return ship.ErrBadRequest.Newf("节点 %s 不存在", arg)
}
