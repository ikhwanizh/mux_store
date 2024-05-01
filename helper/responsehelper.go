package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, data interface{}, statusCode int) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func ResponseError(w http.ResponseWriter, message string, statusCode int) {
	response := map[string]string{"error": message}
	ResponseJson(w, response, statusCode)
	w.WriteHeader(statusCode)
}
