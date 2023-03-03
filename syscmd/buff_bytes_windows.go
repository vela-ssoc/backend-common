package syscmd

import (
	"io"

	"github.com/gorilla/websocket"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func (lb *limitedBuffer) Bytes() []byte {
	buf := lb.bytes()
	if dec, err := simplifiedchinese.GB18030.NewDecoder().Bytes(buf); err == nil {
		return dec
	}
	return buf
}

func (sc *socketConn) Read(p []byte) (int, error) {
	for {
		if sc.rd == nil {
			if _, nr, err := sc.ws.NextReader(); err != nil {
				return 0, err
			} else {
				sc.rd = transform.NewReader(nr, simplifiedchinese.GB18030.NewEncoder())
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
	w, err := sc.ws.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer w.Close()

	return transform.NewWriter(w, simplifiedchinese.GB18030.NewDecoder()).Write(p)
}

func shellEnv() string {
	return "powershell.exe"
}
