package errno

import "github.com/xgfone/ship/v5"

func NodeNotfound(arg string) ship.HTTPServerError {
	return ship.ErrBadRequest.Newf("节点 %s 不存在", arg)
}

func NodeInactive(id int64, inet string) ship.HTTPServerError {
	return ship.ErrBadRequest.Newf("节点 %d(%s) 未激活", id, inet)
}

func NodeOffline(id int64, inet string) ship.HTTPServerError {
	return ship.ErrBadRequest.Newf("节点 %d(%s) 已离线", id, inet)
}

func NodeRemove(id int64, inet string) ship.HTTPServerError {
	return ship.ErrBadRequest.Newf("节点 %d(%s) 已删除", id, inet)
}
