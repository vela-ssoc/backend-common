package netutil

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"sync"

	"github.com/vela-ssoc/backend-common/opurl"
	"github.com/vela-ssoc/backend-common/problem"
)

// Forwarder 代理转发模块
type Forwarder interface {
	Forward(opurl.URLer, http.ResponseWriter, *http.Request)
}

// Forward 新建 forward 代理
func Forward(tran *http.Transport, node string) Forwarder {
	newFn := func() any {
		return &httputil.ReverseProxy{
			Transport: tran,
			ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
				code := http.StatusBadRequest
				ret := &problem.Problem{
					Type:     node,
					Title:    "代理转发错误",
					Status:   code,
					Detail:   err.Error(),
					Instance: r.RequestURI,
				}
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(code)
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
