package core

import "os"

type GcloudConfig struct {
	ProjectID   string
	Location    string
	AccessToken string
}

func NewGCloudConfig() *GcloudConfig {
	return &GcloudConfig{
		ProjectID:   os.Getenv("GCLOUD_PROJECT_ID"),
		Location:    os.Getenv("GCLOUD_LOCATION"),
		AccessToken: os.Getenv("GCLOUD_ACCESS_TOKEN"),
	}
}
