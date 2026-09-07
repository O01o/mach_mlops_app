package rpw

import (
	rpi "mach-mlops/internal/repository/pipeline/interface"
	sc "mach-mlops/internal/schema"
)

type trainAPIRepository struct {
}

func NewTrainAPIRepository() rpi.TrainAPIRepository {
	return &trainAPIRepository{}
}

func (r *trainAPIRepository) RequestParams(params sc.TrainAPIParams) error {
	return nil
}

func (r *trainAPIRepository) RequestExecute() error {
	return nil
}
