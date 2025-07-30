package response

import (
	"encoding/json"
	"mail-phone-auth/internal/entity"
	"net/http"
)

func JSON[T any](w http.ResponseWriter, data *T, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}

func ResponseStatus(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}

func Error(w http.ResponseWriter, msg string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorData := entity.Error{
		Message: msg,
	}

	json.NewEncoder(w).Encode(errorData)
}
