package rpw

import (
	rpi "mach-mlops/internal/repository/pipeline/interface"
	sc "mach-mlops/internal/schema"
)

type trainBatchRepository struct {
}

func NewTrainBatchRepository() rpi.TrainBatchRepository {
	return &trainBatchRepository{}
}

func (r *trainBatchRepository) Execute(params sc.TrainAPIParams) error {
	return nil
}
