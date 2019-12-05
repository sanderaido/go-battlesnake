package snake_test

import (
	"bytes"
	"encoding/json"
	"github.com/sanderaido/go-battlesnake/game"
	"github.com/sanderaido/go-battlesnake/snake"
	"github.com/sanderaido/go-battlesnake/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexResponse(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	snake.HandleIndexRequest(response, request)

	got := response.Body.String()
	want := "This is a Battlesnake participant server"

	if got != want {
		t.Errorf("wanted %q, got %q", want, got)
	}
}

func TestPingResponse(t *testing.T) {

	// NOTE: it's not required in the Battlesnake API that only POST should be allowed here,
	// that's just the request type that the server makes. I just wanted to deny other
	// methods and write a test for it anyway

	var pingTests = []struct {
		method   string
		status   int
	}{
		{http.MethodPost, http.StatusOK},
		{http.MethodGet, http.StatusMethodNotAllowed},
		{http.MethodDelete, http.StatusMethodNotAllowed},
		{http.MethodPut, http.StatusMethodNotAllowed},
		{http.MethodPatch, http.StatusMethodNotAllowed},
	}

	for _, pingTest := range pingTests {
		t.Run("returns HTTP response code 200 only for POST request, otherwise returns 405", func(t *testing.T) {
			request, _ := http.NewRequest(pingTest.method, "/ping", nil)
			response := httptest.NewRecorder()

			snake.HandlePingRequest(response, request)

			got := response.Code
			want := pingTest.status

			if got != want {
				t.Errorf("wanted status %d for %q, got %d", want, pingTest.method, got)
			}
		})
	}
}

var mockPostRequestBody = []byte(`{
  "game": {
    "id": "game-id-string"
  },
  "turn": 4,
  "board": {
    "height": 15,
    "width": 15,
    "food": [
      {
        "x": 1,
        "y": 3
      }
    ],
    "snakes": [
      {
        "id": "snake-id-string",
        "name": "Sneky Snek",
        "health": 90,
        "body": [
          {
            "x": 1,
            "y": 3
          }
        ]
      }
    ]
  },
  "you": {
    "id": "snake-id-string",
    "name": "Sneky Snek",
    "health": 90,
    "body": [
      {
        "x": 1,
        "y": 3
      }
    ]
  }
}`)

func TestHandleStartRequest(t *testing.T) {
	validRequest, _ := http.NewRequest(http.MethodPost, "/start", bytes.NewBuffer(mockPostRequestBody))
	response := httptest.NewRecorder()
	snake.HandleStartRequest(response, validRequest)

	decodedResponse := game.StartResponse{}
	err := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		t.Errorf("bad start request: %v", err)
	}

	color := decodedResponse.Color
	if !strings.HasPrefix(color, "#") || len(color) != 7 {
		t.Errorf("Invalid Color property in start response: %q", color)
	}

	allowedHeadTypes := []string{"beluga", "bendr", "dead", "evil", "fang", "pixel", "regular", "safe", "sand-worm", "shades", "silly", "smile", "tongue"}
	if !util.ContainsString(allowedHeadTypes, decodedResponse.HeadType) {
		t.Errorf("Missing or invalid HeadType property in start response: %q", decodedResponse.HeadType)
	}

	allowedTailTypes := []string{"block-bum", "bolt", "curled", "fat-rattle", "freckled", "hook", "pixel", "regular", "round-bum", "sharp", "skinny", "small-rattle"}
	if !util.ContainsString(allowedTailTypes, decodedResponse.TailType) {
		t.Errorf("Missing or invalid TailType property in start response: %q", decodedResponse.TailType)
	}
}

func TestMoveResponse(t *testing.T) {
	validRequest, _ := http.NewRequest(http.MethodPost, "/move", bytes.NewBuffer(mockPostRequestBody))
	response := httptest.NewRecorder()
	snake.HandleMoveRequest(response, validRequest)
	decodedResponse := game.MoveResponse{}

	err := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		t.Errorf("bad move request: %v", err)
	}

	switch decodedResponse {
	case
		game.MoveResponse{Move:"left"},
		game.MoveResponse{Move:"right"},
		game.MoveResponse{Move:"up"},
		game.MoveResponse{Move:"down"}:
		return
	default:
		t.Errorf("got an invalid move response: %q", decodedResponse)
	}
}

func TestEndResponse(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/end", nil)
	response := httptest.NewRecorder()

	snake.HandleEndRequest(response, request)

	got := response.Code
	want := http.StatusOK

	if got != want {
		t.Errorf("wanted status %d for EndRequest response, got %d", want, got)
	}
}