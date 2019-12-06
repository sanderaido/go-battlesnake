package util

import (
	"encoding/json"
	"errors"
	"github.com/sanderaido/go-battlesnake/game"
	"log"
	"net/http"
)

func RespondJSON(response http.ResponseWriter, data interface{})  {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(data)
	if err == nil {
		_, writeError := response.Write(jsonData)
		if writeError != nil {
			log.Printf("Error writing the response: %v", writeError)
		}
	}
}

func ContainsString(stringArray []string, stringToFind string) bool {
	for _, iteratedString := range stringArray {
		if iteratedString == stringToFind {
			return true
		}
	}
	return false
}

func DecodeMoveRequest(request *http.Request) (game.MoveRequest, error) {
	decodedRequest := game.MoveRequest{}
	err := json.NewDecoder(request.Body).Decode(&decodedRequest)
	if err != nil {
		return game.MoveRequest{}, errors.New("couldn't decode request")
	}
	return decodedRequest, nil
}
