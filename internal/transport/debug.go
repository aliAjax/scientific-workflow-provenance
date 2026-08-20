package transport

import (
	"net/http"
	"net/http/pprof"
)

func pprofHandler(mux *http.ServeMux) {
	mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
}
