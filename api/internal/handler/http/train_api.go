package hh

import (
	"encoding/json"
	core "mach-mlops/internal/core"
	sc "mach-mlops/internal/schema"
	sei "mach-mlops/internal/service/interface"
	"net/http"
)

type apiHandler struct {
	Log *core.Log
	s   sei.TrainAPIService
}

func NewAPIHandler(s sei.TrainAPIService) *apiHandler {
	return &apiHandler{
		Log: core.NewLog(),
		s:   s,
	}
}

func (h *apiHandler) RequestParams(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req sc.TrainAPIParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.s.RequestParams(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *apiHandler) RequestExecute(w http.ResponseWriter, r *http.Request) {
	err := h.s.RequestExecute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
