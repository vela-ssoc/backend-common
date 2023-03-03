package opurl

import (
	"net/http"
	"strconv"
)

var OpPing = URL{method: http.MethodGet, path: "/api/ping", desc: "通用 ping 心跳"}

var BrkJoin = URL{method: http.MethodConnect, path: "/api/broker", desc: "broker 连接 manager 认证"}

var MonJoin = URL{method: http.MethodConnect, path: "/api/minion", desc: "agent(minion) 连接 broker 认证"}

// ------------------------------------
// M: manager  B: broker  A: agent/minion
// rr: request-response 请求响应模式
// ws: websocket 模式
// MA: manager 向 agent 发起
// MB: manager 向 broker 发起
// BA: broker 向 agent 发起
// MArr: manager 向 agent 发起的请求响应请求
// MBws: manager 向 broker 发起的 websocket 双向流
// ...... 按照此规律以此类推
// ------------------------------------

// MArr manager -> agent 请求响应路径
func MArr(bid, mid int64, method, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	sid := strconv.FormatInt(mid, 10)

	return URL{
		host:   host,
		method: method,
		path:   "/api/arr/" + sid + "/" + path,
		query:  query,
		desc:   "manager->agent 请求响应调用",
	}
}

// MBrr manager -> broker 请求响应路径
func MBrr(bid int64, method, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	return URL{
		host:   host,
		method: method,
		path:   "/api/brr/" + path,
		query:  query,
		desc:   "manager->broker 请求响应调用",
	}
}

// MAws manager -> agent 发起的 websocket 请求
func MAws(bid, mid int64, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	sid := strconv.FormatInt(mid, 10)
	return URL{
		scheme: "ws",
		host:   host,
		method: http.MethodGet,
		path:   "/api/aws/" + sid + "/" + path,
		query:  query,
		desc:   "manager->agent websocket 调用",
	}
}

// MBws manager -> broker 发起的 websocket 请求
func MBws(bid int64, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	return URL{
		scheme: "ws",
		host:   host,
		method: http.MethodGet,
		path:   "/api/bws/" + path,
		query:  query,
		desc:   "manager->broker websocket 调用",
	}
}

func BArr(mid, method, path, query string) URL {
	return URL{
		host:   mid,
		method: method,
		path:   "/api/arr/" + path,
		query:  query,
		desc:   "manager->agent 请求响应调用",
	}
}

func BAws(mid, path, query string) URL {
	return URL{
		scheme: "ws",
		host:   mid,
		method: http.MethodGet,
		path:   "/api/aws/" + path,
		query:  query,
		desc:   "broker->agent websocket 调用",
	}
}
