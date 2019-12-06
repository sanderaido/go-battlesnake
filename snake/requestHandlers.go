package snake

import (
	"github.com/sanderaido/go-battlesnake/game"
	"github.com/sanderaido/go-battlesnake/util"
	"log"
	"net/http"
)

func HandleIndexRequest(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	_, writeError := response.Write([]byte("This is a Battlesnake participant server"))
	if writeError != nil {
		log.Printf("Error writing the response: %v", writeError)
	}
}

func HandlePingRequest(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleStartRequest(response http.ResponseWriter, request *http.Request) {
	_, err := util.DecodeMoveRequest(request)
	if err != nil {
		log.Printf("Bad start request: %v", err)
		return
	}

	util.RespondJSON(response, game.StartResponse{
		Color:    "#ff00ff",
		HeadType: "bendr",
		TailType: "pixel",
	})
}

func HandleMoveRequest(response http.ResponseWriter, request *http.Request) {
	_, err := util.DecodeMoveRequest(request)
	if err != nil {
		log.Printf("Bad move request: %v", err)
		return
	}

	util.RespondJSON(response, game.MoveResponse{
		Move: "left",
	})
}

func HandleEndRequest(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(http.StatusOK)
}

