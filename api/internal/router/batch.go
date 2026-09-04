package r

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPBatchHandler() *http.Server {
	r := mux.NewRouter()
	s := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	return s
}
