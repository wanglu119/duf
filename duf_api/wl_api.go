package duf_api

import (
	"encoding/json"
)

func DiskInfoJsonStr() (string, error) {
	m, _, err := mounts()
	if err != nil {
		return "", err
	}

	output, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func DiskMountInfo() ([]Mount, error) {
	m, _, err := mounts()
	if err != nil {
		return m, err
	}
	return m, err
}

