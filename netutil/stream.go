package netutil

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Streamer interface {
	Stream(string, http.Header) (*websocket.Conn, *http.Response, error)
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

func (ss *socketStream) Stream(addr string, header http.Header) (*websocket.Conn, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, res, err := ss.dial.DialContext(ctx, addr, header)
	cancel()

	return conn, res, err
}
