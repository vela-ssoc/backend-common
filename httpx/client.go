package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client HTTP 客户端
type Client struct {
	cli *http.Client // 底层 http.Client
}

// NewClient 新建客户端
func NewClient(trip ...http.RoundTripper) Client {
	cli := http.DefaultClient
	if len(trip) > 0 {
		cli = &http.Client{Transport: trip[0]}
	}

	return Client{
		cli: cli,
	}
}

// Fetch 发送请求
func (c Client) Fetch(ctx context.Context, method string, addr *url.URL, body io.Reader, header http.Header) (*http.Response, error) {
	return c.fetch(ctx, method, addr, body, header)
}

// Do 发送请求
func (c Client) Do(ctx context.Context, method, addr string, body io.Reader, header http.Header) (*http.Response, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	return c.fetch(ctx, method, u, body, header)
}

// DoJSON 发送 JSON 请求
func (c Client) DoJSON(ctx context.Context, method, addr string, body, resp any, header http.Header) error {
	res, err := c.doJSON(ctx, method, addr, body, header)
	if err != nil {
		return err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(resp)
}

// DoJSONTimeout 超时控制发送请求
func (c Client) DoJSONTimeout(timeout time.Duration, method, addr string, body, resp any, header http.Header) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	err := c.DoJSON(ctx, method, addr, body, resp, header)
	cancel()
	return err
}

// NewRequest 构造 http.Request
func (c Client) NewRequest(ctx context.Context, method string, addr *url.URL, body io.Reader, header http.Header) *http.Request {
	return c.newRequest(ctx, method, addr, body, header)
}

// fetchJSON 发送 json 数据，返回原始响应
func (c Client) fetchJSON(ctx context.Context, method string, addr *url.URL, body any, header http.Header) (*http.Response, error) {
	rd, err := c.toJSON(body)
	if err != nil {
		return nil, err
	}

	if header == nil {
		header = make(http.Header, 2)
	}
	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json; charset=UTF-8")

	return c.Fetch(ctx, method, addr, rd, header)
}

func (c Client) doJSON(ctx context.Context, method, addr string, body any, header http.Header) (*http.Response, error) {
	rd, err := c.toJSON(body)
	if err != nil {
		return nil, err
	}
	return c.do(ctx, method, addr, rd, header)
}

func (c Client) do(ctx context.Context, method, addr string, body io.Reader, header http.Header) (*http.Response, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	return c.fetch(ctx, method, u, body, header)
}

// fetch 发送请求
func (c Client) fetch(ctx context.Context, method string, addr *url.URL, body io.Reader, header http.Header) (*http.Response, error) {
	if addr == nil {
		return nil, &net.AddrError{Err: "target url is nil"}
	}

	req := c.NewRequest(ctx, method, addr, body, header)
	res, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	code := res.StatusCode
	if code >= http.StatusOK && code < http.StatusBadRequest { // 200 <= code < 400
		return res, nil
	}

	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()
	buf := make([]byte, 1024)
	n, _ := io.ReadFull(res.Body, buf)
	err = &Error{
		Code: code,
		Body: buf[:n],
	}

	return nil, err
}

// newRequest 构造 http.Request
func (c Client) newRequest(ctx context.Context, method string, addr *url.URL, body io.Reader, header http.Header) *http.Request {
	if method == "" {
		method = http.MethodGet
	}
	if addr == nil {
		addr = &url.URL{Scheme: "about", Opaque: "blank"} // about:blank
	}
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = io.NopCloser(body)
	}

	req := &http.Request{
		Method:     method,
		URL:        addr,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
		Body:       rc,
		Host:       addr.Host,
	}
	if req.Header == nil {
		req.Header = make(http.Header, 2)
	}
	host := req.Header.Get("Host")
	if host == "" {
		host = addr.Host
	}
	if strings.LastIndex(host, ":") > strings.LastIndex(host, "]") {
		host = strings.TrimSuffix(host, ":")
	}
	if body != nil {
		switch v := body.(type) {
		case *bytes.Buffer:
			req.ContentLength = int64(v.Len())
			buf := v.Bytes()
			req.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(buf)
				return io.NopCloser(r), nil
			}
		case *bytes.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return io.NopCloser(&r), nil
			}
		case *strings.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return io.NopCloser(&r), nil
			}
		default:
			// This is where we'd set it to -1 (at least
			// if body != NoBody) to mean unknown, but
			// that broke people during the Go 1.8 testing
			// period. People depend on it being 0 I
			// guess. Maybe retry later. See Issue 18117.
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

// toJSON 转为 JSON
func (Client) toJSON(v any) (io.Reader, error) {
	if v == nil {
		return nil, nil
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		return nil, err
	}
	return buf, nil
}
