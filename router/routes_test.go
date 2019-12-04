package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/sanderaido/go-battlesnake/router"
)

func TestIndex(t *testing.T) {
	t.Run("returns basic welcome message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		router.Index(response, request)

		got := response.Body.String()
		want := "This is a Battlesnake server"

		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})
}