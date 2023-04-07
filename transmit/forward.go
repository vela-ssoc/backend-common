package transmit

import (
	"net/http"
	"net/http/httputil"
	"sync"

	"github.com/vela-ssoc/backend-common/problem"
	"github.com/vela-ssoc/backend-common/transmit/opcode"
)

type Forwarder interface {
	Forward(opcode.URLer, http.ResponseWriter, *http.Request)
}

func NewForward(trip http.RoundTripper, name string) Forwarder {
	newFn := func() any {
		return &httputil.ReverseProxy{
			Transport: trip,
			ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
				code := http.StatusBadRequest
				pd := &problem.Detail{
					Type:     name,
					Title:    "代理转发错误",
					Status:   code,
					Detail:   err.Error(),
					Instance: r.RequestURI,
				}
				_ = pd.JSON(w)
			},
		}
	}

	return &proxy{
		pool: sync.Pool{New: newFn},
	}
}

type proxy struct {
	pool sync.Pool
}

func (p *proxy) Forward(op opcode.URLer, w http.ResponseWriter, r *http.Request) {
	px := p.get()
	defer p.put(px)

	px.Rewrite = func(r *httputil.ProxyRequest) {
		r.Out.URL = op.URL()
		r.SetXForwarded()
	}
	px.ServeHTTP(w, r)
}

func (p *proxy) get() *httputil.ReverseProxy {
	px := p.pool.Get().(*httputil.ReverseProxy)
	return px
}

func (p *proxy) put(px *httputil.ReverseProxy) {
	p.pool.Put(px)
}
