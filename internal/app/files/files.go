package files

import (
	"encoding/json"
	"os"
)

func InitConfig[T any](configFilePath string) (*T, error) {
	
	var config T

	fileData, err := os.ReadFile(configFilePath)

	if err != nil {
		return  nil, err
	}

	err = json.Unmarshal(fileData, &config)

	if err != nil {
		return  nil, err
	}

	return &config, nil
}

