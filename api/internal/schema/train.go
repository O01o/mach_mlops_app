package sc

type TrainAPIParams struct {
	BatchSize int `json:"batch_size"`
	Epoch     int `json:"epoch"`
}

type TrainAPIParamsWorkflows struct {
	Argument TrainAPIParams `json:"argument"`
}
