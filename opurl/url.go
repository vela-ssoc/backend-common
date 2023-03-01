package opurl

import (
	"net/http"
	"net/url"
	"strconv"
)

// URLer 类似于操作码
type URLer interface {
	// Method HTTP 方法
	Method() string

	// String url.URL.String
	String() string

	// URL 组合的 url
	URL() *url.URL

	// Desc 接口说明或备注
	Desc() string
}

type URL struct {
	scheme string
	host   string
	method string
	path   string
	query  string
	desc   string
}

func (u URL) SetQuery(query url.Values) URL { u.query = query.Encode(); return u }
func (u URL) WithQuery(query string) URL    { u.query = query; return u }
func (u URL) String() string                { return u.URL().String() }
func (u URL) IntID(id int64) URL            { u.host = strconv.FormatInt(id, 10); return u }
func (u URL) StrID(id string) URL           { u.host = id; return u }
func (u URL) AsWS() URL                     { u.scheme = "ws"; u.method = http.MethodGet; return u }
func (u URL) SetDesc(str string) URL        { u.desc = str; return u }
func (u URL) Desc() string                  { return u.desc }

func (u URL) Method() string {
	if m := u.method; m != "" {
		return m
	}
	return http.MethodGet
}

func (u URL) URL() *url.URL {
	host, scheme := u.host, u.scheme
	if host == "" {
		host = "default"
	}
	if scheme == "" {
		scheme = "http"
	}
	return &url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     u.path,
		RawQuery: u.query,
	}
}
