package r

import (
	"mach-mlops/internal/core"
	hh "mach-mlops/internal/handler/http"
	sem "mach-mlops/internal/service/impl"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPAPIHandler() *http.Server {
	r := mux.NewRouter()

	gcloud := core.NewGCloudConfig()
	headerHandler := hh.NewHeaderHandler()
	r.Use(headerHandler.CORS)

	apiHandler := hh.NewAPIHandler(sem.NewTrainAPIService(gcloud))
	r.HandleFunc("/", apiHandler.RequestParams).Methods("POST")
	r.HandleFunc("/", apiHandler.RequestExecute).Methods("HEAD")

	s := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	return s
}
