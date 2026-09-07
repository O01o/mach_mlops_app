package rpi

import sc "mach-mlops/internal/schema"

type TrainBatchRepository interface {
	Execute(params sc.TrainAPIParams) error
}
