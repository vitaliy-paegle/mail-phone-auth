package files

import (
	"encoding/json"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
)

func InitConfig[T any](configFilePath string) (*T, error) {

	var config T

	fileData, err := os.ReadFile(configFilePath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileData, &config)

	if err != nil {
		return nil, err
	}

	err = Validate(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

func Validate(data any) error {

	validate := validator.New(validator.WithPrivateFieldValidation())

	err := validate.Struct(data)

	if err != nil {
		log.Println("ConfigValidate", err)
	}

	return err
}
