package main

import (
	"github.com/sanderaido/go-battlesnake/snake"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", snake.HandleIndexRequest)
	http.HandleFunc("/ping", snake.HandlePingRequest)
	http.HandleFunc("/start", snake.HandleStartRequest)
	http.HandleFunc("/move", snake.HandleMoveRequest)
	http.HandleFunc("/end", snake.HandleEndRequest)

	port := "8080"

	log.Printf("Running server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
