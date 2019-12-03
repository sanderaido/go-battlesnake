package main

import (
	"log"
	"net/http"
)

type server struct{}

func main() {
	battlesnakeServer := &server{}
	http.Handle("/", battlesnakeServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (battlesnakeServer *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(`{"message": "hello battlesnake"}`))
}
