package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)

	port := "8080"

	log.Printf("Running server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
