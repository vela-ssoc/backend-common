package netutil

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"sync"

	"github.com/vela-ssoc/backend-common/opurl"
	"github.com/vela-ssoc/backend-common/pubody"
)

// Forwarder 代理转发模块
type Forwarder interface {
	Forward(opurl.URLer, http.ResponseWriter, *http.Request)
}

func Forward(tran *http.Transport, node string) Forwarder {
	newFn := func() any {
		return &httputil.ReverseProxy{
			Transport: tran,
			ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
				code := http.StatusBadGateway
				ret := &pubody.BizError{Code: code, Node: node, Cause: err.Error()}
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(ret)
			},
		}
	}
	pool := &sync.Pool{New: newFn}

	return &httpForward{
		pool: pool,
	}
}

type httpForward struct {
	pool *sync.Pool
}

func (hf *httpForward) Forward(op opurl.URLer, w http.ResponseWriter, r *http.Request) {
	px := hf.get()
	defer hf.put(px)

	px.Rewrite = func(r *httputil.ProxyRequest) {
		r.Out.URL = op.URL()
		r.SetXForwarded()
	}
	px.ServeHTTP(w, r)
}

func (hf *httpForward) get() *httputil.ReverseProxy  { return hf.pool.Get().(*httputil.ReverseProxy) }
func (hf *httpForward) put(p *httputil.ReverseProxy) { hf.pool.Put(p) }
