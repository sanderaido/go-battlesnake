package snake

import (
	"encoding/json"
	"github.com/sanderaido/go-battlesnake/game"
	"github.com/sanderaido/go-battlesnake/util"
	"log"
	"net/http"
)

func IndexResponse(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("This is a Battlesnake participant server"))
}

func PingResponse(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func StartResponse(response http.ResponseWriter, request *http.Request) {
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

func MoveResponse(response http.ResponseWriter, request *http.Request) {
	decoded := game.MoveRequest{}
	err := json.NewDecoder(request.Body).Decode(&decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
		return
	}

	util.RespondJSON(response, game.MoveResponse{
		Move: "left",
	})
}

func EndResponse(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(http.StatusOK)
}