package sem

import (
	"mach-mlops/internal/core"
	rpi "mach-mlops/internal/repository/pipeline/interface"
	rpw "mach-mlops/internal/repository/pipeline/workflows"
	sc "mach-mlops/internal/schema"
	sei "mach-mlops/internal/service/interface"
)

type trainAPIService struct {
	rp rpi.TrainAPIRepository
}

func NewTrainAPIService(gcloud *core.GcloudConfig) sei.TrainAPIService {
	return &trainAPIService{
		rp: rpw.NewTrainAPIRepository(
			gcloud.ProjectID,
			gcloud.Location,
			gcloud.AccessToken,
		),
	}
}

func (t *trainAPIService) RequestParams(params sc.TrainAPIParams) error {
	return t.rp.RequestParams(params)
}

func (t *trainAPIService) RequestExecute() error {
	return t.rp.RequestExecute()
}
