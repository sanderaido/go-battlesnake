package main

import (
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("This is a Battlesnake server"))
}
