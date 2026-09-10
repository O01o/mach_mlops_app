package hh

import (
	"encoding/json"
	core "mach-mlops/internal/core"
	sc "mach-mlops/internal/schema"
	sei "mach-mlops/internal/service/interface"
	"net/http"
)

type batchHandler struct {
	Log *core.Log
	s   sei.TrainBatchService
}

func NewBatchHandler(s sei.TrainBatchService) *batchHandler {
	return &batchHandler{
		Log: core.NewLog(),
		s:   s,
	}
}

func (h *batchHandler) RequestParams(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req sc.TrainAPIParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.s.Execute(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
