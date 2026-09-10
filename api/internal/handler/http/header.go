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
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:18181")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, HEAD, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
