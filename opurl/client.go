package opurl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"

	"github.com/vela-ssoc/backend-common/logback"
)

// NewClient 创建 http client
func NewClient(tran http.RoundTripper, slog logback.Logger) Client {
	return Client{
		cli:  &http.Client{Transport: tran},
		slog: slog,
	}
}

// Client http 客户端
type Client struct {
	cli  *http.Client   // http client
	slog logback.Logger // 日志打印
}

// Fetch 发送请求
func (c Client) Fetch(ctx context.Context, op URLer, header http.Header, body io.Reader) (*http.Response, error) {
	req := c.NewRequest(ctx, op, header, body)
	res, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	code := res.StatusCode
	if code >= http.StatusOK && code < http.StatusMultipleChoices {
		return res, nil
	}
	buf := make([]byte, 1024)
	n, _ := io.ReadFull(res.Body, buf)
	_ = res.Body.Close()
	err = &Error{Code: code, Text: buf[:n]}

	return nil, err
}

// JSON 请求的数据会被 json 序列化后发送，如果 body 为 nil 也可以发送。响应数据会 json 反序列化。
func (c Client) JSON(ctx context.Context, op URLer, header http.Header, body, reply any) error {
	res, err := c.fetchJSON(ctx, op, header, body)
	if err != nil {
		return err
	}
	return json.NewDecoder(res.Body).Decode(reply)
}

// OnewayJSON 发送 json 数据，不关心返回值
func (c Client) OnewayJSON(ctx context.Context, op URLer, header http.Header, body any) error {
	res, err := c.fetchJSON(ctx, op, header, body)
	if err == nil {
		_ = res.Body.Close()
	}
	return nil
}

// Attachment 下载文件接口
func (c Client) Attachment(ctx context.Context, op URLer) (Attachment, error) {
	resp, err := c.Fetch(ctx, op, nil, nil)
	if err != nil {
		return Attachment{}, err
	}
	att := Attachment{rc: resp.Body}
	cd := resp.Header.Get("Content-Disposition")
	if _, params, _ := mime.ParseMediaType(cd); params != nil {
		att.Filename = params["filename"]
		att.Checksum = params["checksum"]
	}
	return att, nil
}

// NewRequest 构造 http.Request
func (Client) NewRequest(ctx context.Context, op URLer, header http.Header, body io.Reader) *http.Request {
	method, dst := op.Method(), op.URL()
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = io.NopCloser(body)
	}
	req := &http.Request{
		Method:     method,
		URL:        dst,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
		Body:       rc,
	}
	if req.Header == nil {
		req.Header = make(http.Header, 4)
	}
	// 设置主机头
	if host := req.Header.Get("Host"); host != "" {
		req.Host = host
	}
	if body != nil {
		switch v := body.(type) {
		case *bytes.Buffer:
			buf := v.Bytes()
			req.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(buf)
				return io.NopCloser(r), nil
			}
		case *bytes.Reader:
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return io.NopCloser(&r), nil
			}
		case *strings.Reader:
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return io.NopCloser(&r), nil
			}
		case io.ReadCloser:
			req.GetBody = func() (io.ReadCloser, error) {
				return v, nil
			}
		default:
			// This is where we'd set it to -1 (at least
			// if body != NoBody) to mean unknown, but
			// that broke people during the Go 1.8 testing
			// period. People depend on it being 0 I
			// guess. Maybe retry later. See Issue 18117.
		}
		if v, yes := body.(interface{ Len() int }); yes {
			req.ContentLength = int64(v.Len())
		}
		// For client requests, Request.ContentLength of 0
		// means either actually 0, or unknown. The only way
		// to explicitly say that the ContentLength is zero is
		// to set the Body to nil. But turns out too much code
		// depends on NewRequest returning a non-nil Body,
		// so we use a well-known ReadCloser variable instead
		// and have the http package also treat that sentinel
		// variable to mean explicitly zero.
		if req.GetBody != nil && req.ContentLength == 0 {
			req.Body = http.NoBody
			req.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
		}
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return req.WithContext(ctx)
}

// fetchJSON 发送的 body 会被 json 序列化
func (c Client) fetchJSON(ctx context.Context, op URLer, header http.Header, body any) (*http.Response, error) {
	rwc := c.toJSON(body)
	if header == nil {
		header = make(http.Header, 4)
	}
	header.Set("Content-Type", "application/json; charset=utf-8")
	header.Set("Accept", "application/json")
	return c.Fetch(ctx, op, header, rwc)
}

func (Client) toJSON(v any) *jsonBody {
	if v == nil {
		return nil
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(v)
	return &jsonBody{err: err, buf: buf}
}

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

type jsonBody struct {
	err error
	buf *bytes.Buffer
}

func (jb *jsonBody) Read(p []byte) (int, error) {
	if jb.err != nil {
		return 0, jb.err
	}
	if jb.buf == nil {
		return 0, io.EOF
	}
	return jb.buf.Read(p)
}

func (jb *jsonBody) Close() error { return nil }

func (jb *jsonBody) Len() int {
	if jb.err != nil || jb.buf == nil {
		return 0
	}
	return jb.buf.Len()
}
