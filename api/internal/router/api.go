package r

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPAPIHandler() *http.Server {
	r := mux.NewRouter()
	s := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	return s
}
