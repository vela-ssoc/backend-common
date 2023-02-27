package netutil

import (
	"context"
	"net"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vela-ssoc/backend-common/opurl"
)

type Streamer interface {
	Stream(opurl.URLer) (*websocket.Conn, error)
}

func Stream(dialFn func(context.Context, string, string) (net.Conn, error)) Streamer {
	dial := &websocket.Dialer{
		NetDialContext:    dialFn,
		HandshakeTimeout:  5 * time.Second,
		ReadBufferSize:    10 * 1024,
		WriteBufferSize:   10 * 1024,
		EnableCompression: true,
	}

	return &socketStream{dial: dial}
}

type socketStream struct {
	dial *websocket.Dialer
}

func (ss *socketStream) Stream(op opurl.URLer) (*websocket.Conn, error) {
	dest := op.String()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, _, err := ss.dial.DialContext(ctx, dest, nil)
	cancel()

	return conn, err
}
