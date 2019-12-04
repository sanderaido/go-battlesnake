package router

import (
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("This is a Battlesnake participant server"))
}

func Ping(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		res.WriteHeader(http.StatusOK)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
	}

}
