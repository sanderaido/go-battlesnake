package util

import (
	"encoding/json"
	"errors"
	"github.com/sanderaido/go-battlesnake/game"
	"net/http"
)

func RespondJSON(response http.ResponseWriter, data interface{})  {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(data)
	if err == nil {
		response.Write(jsonData)
	}
}

func ContainsString(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func DecodeMoveRequest(request *http.Request) (game.MoveRequest, error) {
	decoded := game.MoveRequest{}
	err := json.NewDecoder(request.Body).Decode(&decoded)
	if err != nil {
		return game.MoveRequest{}, errors.New("couldn't decode request")
	}
	return decoded, nil
}
