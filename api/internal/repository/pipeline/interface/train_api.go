package rpi

import sc "mach-mlops/internal/schema"

type TrainAPIRepository interface {
	RequestParams(params sc.TrainAPIParams) error
	RequestExecute() error
}
