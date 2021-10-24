package profile

import (
	"net/http"
	"net/http/pprof"
)

// RegisterProfiler adds profile endpoints to mux.
func RegisterProfiler(mux *http.ServeMux) {
	mux.HandleFunc("/debug/profile/", pprof.Index)
	mux.HandleFunc("/debug/profile/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/profile/profile", pprof.Profile)
	mux.HandleFunc("/debug/profile/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/profile/trace", pprof.Trace)
}
