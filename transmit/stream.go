package transmit

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vela-ssoc/backend-common/transmit/opcode"
)

type Streamer interface {
	Stream(opcode.URLer, http.Header) (*websocket.Conn, *http.Response, error)
}

func NewStream(dialFn func(context.Context, string, string) (net.Conn, error)) Streamer {
	dial := &websocket.Dialer{
		NetDialContext:    dialFn,
		HandshakeTimeout:  5 * time.Second,
		ReadBufferSize:    4 * 1024,
		WriteBufferSize:   4 * 1024,
		EnableCompression: true,
	}

	return &socketStream{dial: dial}
}

type socketStream struct {
	dial *websocket.Dialer
}

func (ss *socketStream) Stream(op opcode.URLer, header http.Header) (*websocket.Conn, *http.Response, error) {
	addr := op.URL().String()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, res, err := ss.dial.DialContext(ctx, addr, header)
	cancel()

	return conn, res, err
}
