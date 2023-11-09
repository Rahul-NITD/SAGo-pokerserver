package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rahul-NITD/SAGo/PokerServer/pokerhttpserver"
)

func TestIntegration(t *testing.T) {
	t.Run("Integration Test", func(t *testing.T) {
		server := pokerhttpserver.PlayerServer{
			Storage: pokerhttpserver.NewInMemoryStorage(),
		}
		req, _ := http.NewRequest(http.MethodPost, "/players/Raaa", nil)
		res := httptest.NewRecorder()

		// make post request 3 times
		server.ServeHTTP(res, req)
		server.ServeHTTP(res, req)
		server.ServeHTTP(res, req)

		req, _ = http.NewRequest(http.MethodGet, "/players/Raaa", nil)
		res = httptest.NewRecorder()
		server.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("Status %v != %v", res.Code, http.StatusOK)
		}

		if res.Body.String() != "3" {
			t.Errorf("Received %v, want %v", res.Body.String(), "3")
		}

	})
}
