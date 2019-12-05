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

func ContainsString(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}