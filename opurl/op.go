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

func MMrr(bid, mid int64, method, path, query string) URL {
	host := strconv.FormatInt(bid, 10)
	sid := strconv.FormatInt(mid, 10)

	return URL{
		host:   host,
		method: method,
		path:   "/api/mrr/" + sid + "/" + path,
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

func BMrr(mid, method, path, query string) URL {
	return URL{
		host:   mid,
		method: method,
		path:   "/api/mrr/" + path,
		query:  query,
	}
}

func BMws(mid, path, query string) URL {
	return URL{
		scheme: "ws",
		host:   mid,
		method: http.MethodGet,
		path:   "/api/mws/" + path,
		query:  query,
	}
}

func MMws(bid, mid int64, path, query string) URL {
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
