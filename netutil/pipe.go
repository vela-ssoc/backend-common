package netutil

import "github.com/gorilla/websocket"

func Pipe(fore, back *websocket.Conn) {
	go func() {
		for {
			mt, p, err := fore.ReadMessage()
			if err == nil {
				err = back.WriteMessage(mt, p)
			}
			if err != nil {
				_ = back.Close()
				_ = fore.Close()
				break
			}
		}
	}()

	for {
		mt, p, err := back.ReadMessage()
		if err == nil {
			err = fore.WriteMessage(mt, p)
		}
		if err != nil {
			_ = fore.Close()
			_ = back.Close()
			break
		}
	}
}
