package sei

import sc "mach-mlops/internal/schema"

type TrainBatchService interface {
	Execute(params sc.TrainAPIParams) error
}
