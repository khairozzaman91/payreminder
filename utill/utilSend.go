package utill

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data any, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func SendError(w http.ResponseWriter, statusCode int, msg string) {
	response := map[string]string{
		"error": msg,
	}
	SendData(w, response, statusCode)
}
