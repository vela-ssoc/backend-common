package pubrr

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"sync"

	"github.com/vela-ssoc/backend-common/opurl"
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
				ret := &ErrorResult{Code: code, Node: node, Cause: err.Error()}
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				_ = json.NewEncoder(w).Encode(ret)
			},
		}
	}
	pool := &sync.Pool{New: newFn}

	return &forward{
		pool: pool,
	}
}

type forward struct {
	pool *sync.Pool
}

func (fd *forward) Forward(op opurl.URLer, w http.ResponseWriter, r *http.Request) {
	px := fd.get()
	defer fd.put(px)

	px.Rewrite = func(r *httputil.ProxyRequest) {
		r.Out.URL = op.URL()
		r.SetXForwarded()
	}
	px.ServeHTTP(w, r)
}

func (fd *forward) get() *httputil.ReverseProxy  { return fd.pool.Get().(*httputil.ReverseProxy) }
func (fd *forward) put(p *httputil.ReverseProxy) { fd.pool.Put(p) }
