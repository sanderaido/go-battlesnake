package router

import (
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("This is a Battlesnake participant server"))
}

func Ping(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}
