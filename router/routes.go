package router

import (
	"encoding/json"
	"github.com/sanderaido/go-battlesnake/game"
	"github.com/sanderaido/go-battlesnake/util"
	"log"
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

func Start(response http.ResponseWriter, request *http.Request) {
	decoded := game.MoveRequest{}
	err := json.NewDecoder(request.Body).Decode(&decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
		return
	}

	util.RespondJSON(response, game.StartResponse{
		Color: "#ff00ff",
		HeadType: "bendr",
		TailType: "pixel",
	})
}
