package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error 返回错误
type Error struct {
	Code int    `json:"code"` // http 返回的 状态码
	Body []byte `json:"body"` // http body，最多 1024 字节
}

func (e *Error) Error() string {
	return fmt.Sprintf("httpx client response status %d, message is: %s", e.Code, e.Body)
}

func (e *Error) NotAcceptable() bool {
	return e.Code == http.StatusNotAcceptable
}

func (e *Error) JSON(v any) error {
	return json.Unmarshal(e.Body, v)
}
