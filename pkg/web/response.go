package web

import (
	"encoding/json"
	"net/http"
)

type ResponseProduct struct {
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Error      bool   `json:"error"`
	StatusCode int    `json:"status_code"`
}

func NewResponseProduct() ResponseProduct {
	return ResponseProduct{}
}

func SendResponse(w http.ResponseWriter, message string, data any, err bool, statusCode int) {
	resposne := &ResponseProduct{
		Message:    message,
		Data:       data,
		Error:      err,
		StatusCode: statusCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resposne)
}
