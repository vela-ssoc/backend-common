package httpclient

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code int
	Text []byte
}

func (e *Error) Error() string {
	return fmt.Sprintf("http response status %d, message is: %s", e.Code, e.Text)
}

func (e *Error) NotAcceptable() bool {
	return e.Code == http.StatusNotAcceptable
}
