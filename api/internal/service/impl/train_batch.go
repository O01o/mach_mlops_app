package sem

import (
	sc "mach-mlops/internal/schema"
	sei "mach-mlops/internal/service/interface"
)

type trainBatchService struct {
}

func NewTrainBatchService() sei.TrainBatchService {
	return &trainBatchService{}
}

func (t *trainBatchService) Execute(params sc.TrainAPIParams) error {
	return nil
}
