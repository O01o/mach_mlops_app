package sem

import (
	sc "mach-mlops/internal/schema"
	sei "mach-mlops/internal/service/interface"
)

type trainAPIService struct {
}

func NewTrainAPIService() sei.TrainAPIService {
	return &trainAPIService{}
}

func (t *trainAPIService) RequestParams(params sc.TrainAPIParams) error {
	return nil
}

func (t *trainAPIService) RequestExecute() error {
	return nil
}
