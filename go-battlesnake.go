package main

import (
	"github.com/sanderaido/go-battlesnake/snake"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", snake.IndexResponse)
	http.HandleFunc("/ping", snake.PingResponse)
	http.HandleFunc("/start", snake.StartResponse)
	http.HandleFunc("/move", snake.MoveResponse)
	http.HandleFunc("/end", snake.EndResponse)

	port := "8080"

	log.Printf("Running server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
