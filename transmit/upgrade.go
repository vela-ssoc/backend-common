package transmit

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vela-ssoc/backend-common/problem"
)

func Upgrade(node string) websocket.Upgrader {
	errorFn := func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		pd := &problem.Detail{
			Type:     node,
			Title:    "websocket 协议升级错误",
			Status:   status,
			Detail:   reason.Error(),
			Instance: r.RequestURI,
		}
		_ = pd.JSON(w)
	}

	return websocket.Upgrader{
		HandshakeTimeout:  5 * time.Second,
		ReadBufferSize:    4 * 1024,
		WriteBufferSize:   4 * 1024,
		Error:             errorFn,
		CheckOrigin:       func(*http.Request) bool { return true },
		EnableCompression: true,
	}
}
