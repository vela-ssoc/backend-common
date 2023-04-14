package opcode

import (
	"net/http"
	"net/url"
	"strconv"
)

type URLer interface {
	Method() string
	Path() string
	URL() *url.URL
	AsWS() URLer
	SetID(string) URLer
	SetIntID(int64) URLer
	SetQuery(string) URLer
}

type opURL struct {
	method string
	scheme string
	host   string
	path   string
	query  string
	desc   string
}

func (op opURL) Method() string {
	method := op.method
	if method == "" {
		method = http.MethodGet
	}
	return method
}

func (op opURL) Path() string {
	path := op.path
	if path == "" {
		path = "/"
	}
	return path
}

func (op opURL) SetID(id string) URLer {
	op.host = id
	return op
}

func (op opURL) SetIntID(id int64) URLer {
	op.host = strconv.FormatInt(id, 10)
	return op
}

func (op opURL) URL() *url.URL {
	scheme, host := op.scheme, op.host
	if scheme == "" {
		scheme = "http"
	}
	if host == "" {
		// 如果为空，发送请求时会校验不通过。
		host = "none"
	}

	return &url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     op.path,
		RawQuery: op.query,
	}
}

func (op opURL) AsWS() URLer {
	op.scheme = "ws"
	return op
}

func (op opURL) SetQuery(q string) URLer {
	op.query = q
	return op
}
