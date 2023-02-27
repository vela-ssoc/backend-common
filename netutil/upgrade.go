package netutil

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vela-ssoc/backend-common/pubody"
)

func Upgrade(node string) websocket.Upgrader {
	errorFn := func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		ret := &pubody.BizError{Code: status, Node: node, Cause: reason.Error()}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ret)
	}

	return websocket.Upgrader{
		HandshakeTimeout:  10 * time.Second,
		ReadBufferSize:    10 * 1024,
		WriteBufferSize:   10 * 1024,
		Error:             errorFn,
		CheckOrigin:       func(*http.Request) bool { return true },
		EnableCompression: true,
	}
}
