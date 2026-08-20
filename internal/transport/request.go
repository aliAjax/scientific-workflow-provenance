package transport

import (
	"net/http"
	"strings"
	"time"
)

type RequestInfo struct {
	ID      string
	Method  string
	Path    string
	Started time.Time
}

func info(r *http.Request) RequestInfo {
	return RequestInfo{ID: r.Header.Get("X-Request-ID"), Method: r.Method, Path: r.URL.Path, Started: time.Now()}
}
func isJSON(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Content-Type"), "application/json")
}
func acceptsJSON(r *http.Request) bool {
	v := r.Header.Get("Accept")
	return v == "" || strings.Contains(v, "json")
}
