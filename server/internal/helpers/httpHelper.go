package httpHelper

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response[T any] struct {
	Data    *T
	Message string
	Success bool
}

func SendResponse[T any](w http.ResponseWriter, data *T, success bool, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := Response[T]{
		Data:    data,
		Success: success,
		Message: message,
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		slog.Error("Failed to serialize response", err)
		SendResponse[interface{}](w, nil, false, "Failed to serialize repsonse", http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}
