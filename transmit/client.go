package transmit

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"time"

	"github.com/vela-ssoc/backend-common/httpx"
	"github.com/vela-ssoc/backend-common/transmit/opcode"
)

func NewClient(trip http.RoundTripper) Client {
	cli := httpx.NewClient(trip)
	return Client{cli: cli}
}

type Client struct {
	cli httpx.Client
}

func (c Client) Fetch(ctx context.Context, op opcode.URLer, rd io.Reader, header http.Header) (*http.Response, error) {
	return c.fetch(ctx, op, rd, header)
}

func (c Client) JSON(ctx context.Context, op opcode.URLer, body, resp any) error {
	res, err := c.fetchJSON(ctx, op, body, nil)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(resp)
}

func (c Client) OnewayJSON(ctx context.Context, op opcode.URLer, body any) error {
	res, err := c.fetchJSON(ctx, op, body, nil)
	if err == nil {
		return res.Body.Close()
	}
	return err
}

// Attachment 下载文件/附件
func (c Client) Attachment(ctx context.Context, op opcode.URLer) (Attachment, error) {
	resp, err := c.fetch(ctx, op, nil, nil)
	if err != nil {
		return Attachment{}, err
	}
	att := Attachment{code: resp.StatusCode, rc: resp.Body}
	cd := resp.Header.Get("Content-Disposition")
	if _, params, _ := mime.ParseMediaType(cd); params != nil {
		att.Filename = params["filename"]
		att.Checksum = params["checksum"]
	}
	return att, nil
}

func (c Client) NewRequest(ctx context.Context, op opcode.URLer, body io.Reader, header http.Header) *http.Request {
	method, addr := op.Method(), op.URL()
	return c.cli.NewRequest(ctx, method, addr, body, header)
}

func (c Client) fetchJSON(ctx context.Context, op opcode.URLer, body any, header http.Header) (*http.Response, error) {
	rd, err := c.toJSON(body)
	if err != nil {
		return nil, err
	}
	if header == nil {
		header = make(http.Header, 2)
	}
	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json; charset=UTF-8")

	return c.fetch(ctx, op, rd, header)
}

func (c Client) fetch(ctx context.Context, op opcode.URLer, rd io.Reader, header http.Header) (*http.Response, error) {
	method, addr := op.Method(), op.URL()
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
	}

	return c.cli.Fetch(ctx, method, addr, rd, header)
}

func (c Client) toJSON(v any) (io.Reader, error) {
	if v == nil {
		return nil, nil
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		return nil, err
	}
	return buf, nil
}
