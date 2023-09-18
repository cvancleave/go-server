package utils

import (
	"encoding/json"
	"go-server/internal/models"
	"net/http"
)

// RespondJson - write a json response
func RespondJson(w http.ResponseWriter, code int, data any) {
	body, _ := json.Marshal(data)
	w.WriteHeader(code)
	w.Write(body)
}

// RespondError - write an error response
func RespondError(w http.ResponseWriter, code int, message string) {
	response := models.ErrorResponse{
		ErrorCode:    code,
		ErrorMessage: message,
	}
	body, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(body)
}
