package opurl

import "net/http"

var (
	OpPing        = URL{method: http.MethodGet, path: "/api/ping"}
	OpIntobSyscmd = URL{method: http.MethodGet, path: "/api/intob/syscmd"}
)
