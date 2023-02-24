package opurl

import (
	"net/http"
	"strconv"
)

var (
	OpPing        = URL{method: http.MethodGet, path: "/api/ping"}
	OpIntobSyscmd = URL{method: http.MethodGet, path: "/api/intob/syscmd"}
	OpIntomSyscmd = URL{method: http.MethodGet, path: "/api/intom/syscmd"}
)

var MonFS = URL{method: http.MethodGet, path: "/api/fs"}

var BrkFS = URL{method: http.MethodGet, path: "/api/intob/fs"}

func MgtIntom(method string, bid, mid int64, path string) URL {
	p := "/api/intom/" + strconv.FormatInt(mid, 10) + "/" + path
	return URL{
		id:     strconv.FormatInt(bid, 10),
		method: method,
		path:   p,
	}
}

func BrkIntom(method string, mid string, path string) URL {
	return URL{
		id:     mid,
		method: method,
		path:   "/api/intom/" + path,
	}
}
