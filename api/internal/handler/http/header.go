package hh

import (
	core "mach-mlops/internal/core"
	"net/http"
)

type headerHandler struct {
	Log *core.Log
}

func NewHeaderHandler() *headerHandler {
	return &headerHandler{
		Log: core.NewLog(),
	}
}

func (h *headerHandler) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// reflect the request Origin since port-forwarding/tunnels change the origin seen by the browser
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Add("Access-Control-Allow-Origin", origin)
			w.Header().Add("Vary", "Origin")
		}
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, HEAD, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
