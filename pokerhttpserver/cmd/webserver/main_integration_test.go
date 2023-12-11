package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Rahul-NITD/SAGo/PokerServer/pokerhttpserver"
)

func TestIntegration(t *testing.T) {
	server := pokerhttpserver.NewPlayerServer(pokerhttpserver.NewInMemoryStorage())
	req, _ := http.NewRequest(http.MethodPost, "/players/Raaa", nil)
	res := httptest.NewRecorder()

	// make post request 3 times
	server.ServeHTTP(res, req)
	server.ServeHTTP(res, req)
	server.ServeHTTP(res, req)
	t.Run("Integration Test /players/", func(t *testing.T) {
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

	t.Run("Integration Test /league", func(t *testing.T) {
		req, _ = http.NewRequest(http.MethodGet, "/league", nil)
		res = httptest.NewRecorder()
		server.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("Status %v != %v", res.Code, http.StatusOK)
		}

		if res.Result().Header.Get("content-type") != "application/json" {
			t.Error("content type not set to application/json")
		}

		var got []pokerhttpserver.Player

		err := json.NewDecoder(res.Body).Decode(&got)

		if err != nil {
			t.Error(err.Error())
		}

		want := []pokerhttpserver.Player{
			{
				Name: "Raaa",
				Wins: 3,
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

}
