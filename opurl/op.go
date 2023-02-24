package opurl

import (
	"net/http"
	"strconv"
)

var OpPing = URL{method: http.MethodGet, path: "/api/ping"}

var BrkJoin = URL{method: http.MethodConnect, path: "/api/broker"}

var MonJoin = URL{method: http.MethodConnect, path: "/api/minion"}

func BIntom(mid, method, path, query string) URL {
	return URL{
		host:   mid,
		method: method,
		path:   "/api/intom/" + path,
		query:  query,
	}
}

func MIntom(bid, mid int64, method, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	sid := strconv.FormatInt(mid, 10)

	return URL{
		host:   host,
		method: method,
		path:   "/api/intom/" + sid + "/" + path,
		query:  query,
	}
}

func Intob(bid int64, method, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	return URL{
		host:   host,
		method: method,
		path:   "/api/intob/" + path,
		query:  query,
	}
}
