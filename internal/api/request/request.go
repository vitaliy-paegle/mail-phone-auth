package request

import (
	"encoding/json"
	"io"
)


func DecodeBody[T any](body io.ReadCloser) (T, error) {
	var data T
	
	err := json.NewDecoder(body).Decode(&data)

	if err != nil {
		return data, err
	} else {
		return data, nil
	}

}