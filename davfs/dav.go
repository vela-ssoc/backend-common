package davfs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/webdav"
)

func FS(dir, prefix string) http.Handler {
	prefix = strings.TrimSuffix(prefix, "/")
	h := &webdav.Handler{
		Prefix:     prefix,
		FileSystem: webdav.Dir(dir),
		LockSystem: webdav.NewMemLS(),
	}
	return &davfs{h: h}
}

type davfs struct {
	h *webdav.Handler
}

func (dav *davfs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if link := r.Header.Get("Link"); link != "" {
		dav.h.Prefix = link
	}
	if r.Method == http.MethodGet && dav.serveHTTP(w, r) {
		return
	}
	dav.h.ServeHTTP(w, r)
}

func (dav *davfs) serveHTTP(w http.ResponseWriter, r *http.Request) bool {
	ctx, path := r.Context(), r.URL.Path
	prefix, fs := dav.h.Prefix, dav.h.FileSystem
	name := strings.TrimPrefix(path, prefix)
	if name == "" {
		name = "/"
	}
	file, err := fs.OpenFile(ctx, name, os.O_RDONLY, 0)
	if err != nil {
		return false
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	if stat, err := file.Stat(); err != nil || !stat.IsDir() {
		return false
	}
	if link := r.Header.Get("Link"); link != "" {
		if !strings.HasSuffix(link, "/") {
			w.Header().Set("Location", link+"/")
			w.WriteHeader(http.StatusMovedPermanently)
			return true
		}
	}

	infos, err := file.Readdir(4096)
	if err != nil && err != io.EOF {
		return false
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(htmlHead)
	for _, info := range infos {
		filename := info.Name()
		if info.IsDir() {
			filename += "/"
		}
		mtime := info.ModTime().In(time.Local)
		_, _ = fmt.Fprintf(w, htmlTbl, filename, filename, info.Mode(), mtime.Format(time.RFC3339))
	}
	_, _ = w.Write(htmlTail)

	return true
}

var (
	htmlHead = []byte(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="color-scheme" content="light dark">
    <title>WebDAV</title>
</head>
<body>
<pre style="word-wrap: break-word; white-space: pre-wrap;">
    <table>
`)

	htmlTbl = `
    <tr>
    <td><a href="%s">%s</a></td>
	<td>%s</td>
	<td>%s</td>
    </tr>

`

	htmlTail = []byte(`
    </table>
</pre>
</body>
</html>
`)
)
