package util

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(response http.ResponseWriter, data interface{})  {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(data)
	if err == nil {
		response.Write(jsonData)
	}
}
