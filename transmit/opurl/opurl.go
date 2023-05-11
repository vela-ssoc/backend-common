package opurl

import (
	"net/http"
	"net/url"
)

const HeaderXNodeID = "X-Node-ID"

type URLer interface {
	URL() *url.URL
	Method() string
}

type opURL struct {
	scheme string
	host   string
	path   string
	query  string
	method string
	nodeID string
}

func (u opURL) URL() *url.URL {
	if u.scheme == "" {
		u.scheme = "http"
	}
	if u.host == "" {
		u.host = "soc"
	}
	if u.path == "" {
		u.path = "/"
	}
	return &url.URL{
		Scheme:   u.scheme,
		Host:     u.host,
		Path:     u.path,
		RawQuery: u.query,
	}
}

func (u opURL) Method() string {
	if u.method == "" {
		return http.MethodGet
	}
	return u.method
}
