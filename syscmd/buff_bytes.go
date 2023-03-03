//go:build !windows

package syscmd

import (
	"io"
	"os"

	"github.com/gorilla/websocket"
)

func (lb *limitedBuffer) Bytes() []byte {
	return lb.bytes()
}

func (sc *socketConn) Read(p []byte) (int, error) {
	for {
		if sc.rd == nil {
			if _, nr, err := sc.ws.NextReader(); err != nil {
				return 0, err
			} else {
				sc.rd = nr
			}
		}
		n, err := sc.rd.Read(p)
		if err == io.EOF {
			sc.rd = nil
			continue
		}
		return n, err
	}
}

func (sc *socketConn) Write(p []byte) (int, error) {
	n := len(p)
	err := sc.ws.WriteMessage(websocket.TextMessage, p)
	return n, err
}

func shellEnv() string {
	name := os.Getenv("SHELL")
	if name == "" {
		name = "/bin/bash"
	}
	return name
}
