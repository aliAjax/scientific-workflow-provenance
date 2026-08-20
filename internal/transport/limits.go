package transport

import (
	"net/http"
	"strconv"
)

func limitBody(r *http.Request, max int64) { r.Body = http.MaxBytesReader(nil, r.Body, max) }
func maxBytes(r *http.Request, def int64) int64 {
	if v, e := strconv.ParseInt(r.Header.Get("X-Max-Bytes"), 10, 64); e == nil && v > 0 {
		return v
	}
	return def
}
