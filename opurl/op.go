package opurl

import (
	"net/http"
	"strconv"
)

var OpPing = URL{method: http.MethodGet, path: "/api/ping"}

var BrkJoin = URL{method: http.MethodConnect, path: "/api/broker"}

var MonJoin = URL{method: http.MethodConnect, path: "/api/minion"}

func MArr(bid, mid int64, method, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	sid := strconv.FormatInt(mid, 10)

	return URL{
		host:   host,
		method: method,
		path:   "/api/arr/" + sid + "/" + path,
		query:  query,
	}
}

func MBrr(bid int64, method, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	return URL{
		host:   host,
		method: method,
		path:   "/api/brr/" + path,
		query:  query,
	}
}

func BArr(mid, method, path, query string) URL {
	return URL{
		host:   mid,
		method: method,
		path:   "/api/arr/" + path,
		query:  query,
	}
}

func BAws(mid, path, query string) URL {
	return URL{
		scheme: "ws",
		host:   mid,
		method: http.MethodGet,
		path:   "/api/mws/" + path,
		query:  query,
	}
}

func MAws(bid, mid int64, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	sid := strconv.FormatInt(mid, 10)
	return URL{
		scheme: "ws",
		host:   host,
		method: http.MethodGet,
		path:   "/api/mws/" + sid + "/" + path,
		query:  query,
	}
}

func MBws(bid int64, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	return URL{
		scheme: "ws",
		host:   host,
		method: http.MethodGet,
		path:   "/api/bws/" + path,
		query:  query,
	}
}
