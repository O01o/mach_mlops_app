package r

import (
	h "mach-mlops/internal/handler"
	sem "mach-mlops/internal/service/impl"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPAPIHandler() *http.Server {
	r := mux.NewRouter()

	hh := h.NewAPIHandler(sem.NewTrainAPIService())
	r.HandleFunc("/", hh.RequestParams).Methods("POST")
	r.HandleFunc("/", hh.RequestExecute).Methods("HEAD")

	s := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	return s
}
