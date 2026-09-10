package sc

type TrainAPIParams struct {
	BatchSize int `json:"batch_size"`
	Epoch     int `json:"epoch"`
}

type TrainAPIParamsWorkflowsArgument struct {
	Message string `json:"message"`
}

type TrainAPIParamsWorkflows struct {
	Argument string `json:"argument"`
}
