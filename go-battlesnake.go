package main

import (
	"log"
	"net/http"
	"github.com/sanderaido/go-battlesnake/router"
)

func main() {
	http.HandleFunc("/", router.Index)
	http.HandleFunc("/ping", router.Ping)

	port := "8080"

	log.Printf("Running server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
