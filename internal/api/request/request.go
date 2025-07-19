package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
)

func DecodeBody[T any](body io.ReadCloser) (T, error) {
	var data T

	err := json.NewDecoder(body).Decode(&data)
	if err != nil {
		return data, err
	}

	err = Validate(data)
	if err != nil {
		return data, err
	} else {
		return data, nil
	}

}

func Validate(data any) error {

	validate := validator.New(validator.WithPrivateFieldValidation())

	err := validate.Struct(data)

	if err != nil {
		log.Println("RequestValidate", err)
	}

	return err
}
