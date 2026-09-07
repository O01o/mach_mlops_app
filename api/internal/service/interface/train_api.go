package sei

import sc "mach-mlops/internal/schema"

type TrainAPIService interface {
	RequestParams(params sc.TrainAPIParams) error
	RequestExecute() error
}
