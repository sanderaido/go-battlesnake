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
		want := "This is a Battlesnake participant server"

		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})
}

func TestPing(t *testing.T) {

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

	t.Run("returns HTTP response code 200 only for POST request, otherwise returns 405", func(t *testing.T) {
		for _, pingTest := range pingTests {
			validRequest, _ := http.NewRequest(pingTest.method, "/ping", nil)
			response := httptest.NewRecorder()

			router.Ping(response, validRequest)

			got := response.Code
			want := pingTest.status

			if got != want {
				t.Errorf("wanted status %d for %q, got %d", want, pingTest.method, got)
			}
		}
	})
}