package rpw

import (
	"bytes"
	"encoding/json"
	"fmt"
	rpi "mach-mlops/internal/repository/pipeline/interface"
	sc "mach-mlops/internal/schema"
	"net/http"
)

type trainAPIRepository struct {
	projectID   string
	location    string
	accessToken string
}

func NewTrainAPIRepository(projectID, location, accessToken string) rpi.TrainAPIRepository {
	return &trainAPIRepository{
		projectID:   projectID,
		location:    location,
		accessToken: accessToken,
	}
}

func (r *trainAPIRepository) RequestParams(params sc.TrainAPIParams) error {
	url := fmt.Sprintf(
		"https://workflowexecutions.googleapis.com/v1/projects/%s/locations/%s/workflows/%s/executions",
		r.projectID,
		r.location,
		"mlops-publish-train-params",
	)
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}
	argument, err := json.Marshal(sc.TrainAPIParamsWorkflowsArgument{Message: string(body)})
	if err != nil {
		return err
	}
	body, err = json.MarshalIndent(sc.TrainAPIParamsWorkflows{Argument: string(argument)}, "", "  ")
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+r.accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s", resp.Status)
	}
	return nil
}

func (r *trainAPIRepository) RequestExecute() error {
	url := fmt.Sprintf(
		"https://workflowexecutions.googleapis.com/v1/projects/%s/locations/%s/workflows/%s/executions",
		r.projectID,
		r.location,
		"mlops-train-batch",
	)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+r.accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s", resp.Status)
	}
	return nil
}
