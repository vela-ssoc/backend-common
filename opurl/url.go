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
	id     string
	method string
	path   string
	query  string
}

func (u URL) SetQuery(query url.Values) URL { u.query = query.Encode(); return u }
func (u URL) WithQuery(query string) URL    { u.query = query; return u }
func (u URL) String() string                { return u.URL().String() }
func (u URL) SetID(id int64) URL            { u.id = strconv.FormatInt(id, 10); return u }
func (u URL) WithID(id string) URL          { u.id = id; return u }

func (u URL) Method() string {
	if m := u.method; m != "" {
		return m
	}
	return http.MethodGet
}

func (u URL) URL() *url.URL {
	return &url.URL{
		Scheme:   "http",
		Host:     u.id,
		Path:     u.path,
		RawQuery: u.query,
	}
}
