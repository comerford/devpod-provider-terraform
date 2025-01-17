package options

import (
	"fmt"
	"os"
)

const (
	REGION            = "REGION"
	LAUNCH_TYPE       = "LAUNCH_TYPE"
	IMAGE_URL         = "IMAGE_URL"
	MEM_ALLOCATION    = "MEM_ALLOCATION"
	TERRAFORM_PROJECT = "TERRAFORM_PROJECT"
)

type Options struct {
	ImageUrl        string
	MemAllocationGB string
	ClusterFolder   string
	ClusterID       string
	LaunchType      string
	Zone            string
}

func ConfigFromEnv() (Options, error) {
	return Options{
		LaunchType:       os.Getenv(LAUNCH_TYPE),
		ImageUrl:         os.Getenv(IMAGE_URL),
		MemAllocationGB:  os.Getenv(MEM_ALLOCATION),
		Zone:             os.Getenv(REGION),
	}, nil
}

func FromEnv() (*Options, error) {
	retOptions := &Options{}

	var err error

	retOptions.ClusterID, err = FromEnvOrError("CLUSTER_ID")
	if err != nil {
		return nil, err
	}
	// prefix with devpod-
	retOptions.ClusterID = "devpod-" + retOptions.ClusterID

	retOptions.ClusterFolder, err = FromEnvOrError("MACHINE_FOLDER")
	if err != nil {
		return nil, err
	}

	retOptions.LaunchType, err = FromEnvOrError("LAUNCH_TYPE")
	if err != nil {
		return nil, err
	}

	retOptions.ImageUrl, err = FromEnvOrError("IMAGE_URL")
	if err != nil {
		return nil, err
	}

	retOptions.MemAllocationGB, err = FromEnvOrError("MEM_ALLOCATION")
	if err != nil {
		return nil, err
	}

	retOptions.Zone, err = FromEnvOrError("REGION")
	if err != nil {
		return nil, err
	}

	return retOptions, nil
}

func FromEnvOrError(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf(
			"couldn't find option %s in environment, please make sure %s is defined",
			name,
			name,
		)
	}

	return val, nil
}
