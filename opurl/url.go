package opurl

import (
	"net/http"
	"net/url"
	"strconv"
)

type URLer interface {
	Method() string
	String() string
	URL() *url.URL
}

type URL struct {
	scheme string
	host   string
	method string
	path   string
	query  string
}

func (u URL) SetQuery(query url.Values) URL { u.query = query.Encode(); return u }
func (u URL) WithQuery(query string) URL    { u.query = query; return u }
func (u URL) String() string                { return u.URL().String() }
func (u URL) IntID(id int64) URL            { u.host = strconv.FormatInt(id, 10); return u }
func (u URL) StrID(id string) URL           { u.host = id; return u }
func (u URL) AsWS() URL                     { u.scheme = "ws"; return u }

func (u URL) Method() string {
	if m := u.method; m != "" {
		return m
	}
	return http.MethodGet
}

func (u URL) URL() *url.URL {
	host := u.host
	if host == "" {
		host = "default"
	}
	scheme := u.scheme
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
