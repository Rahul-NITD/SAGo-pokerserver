package pokerhttpserver_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rahul-NITD/SAGo/PokerServer/pokerhttpserver"
)

type TestStorage struct {
	store map[string]int
	Calls []string
}

func (ts *TestStorage) GetPlayerScore(player string) (int, error) {
	score, ok := ts.store[player]
	if !ok {
		return 0, errors.New("Player Not Found Error")
	}
	return score, nil
}

func (ts *TestStorage) RecordWin(player string) error {
	ts.Calls = append(ts.Calls, player)
	return nil
}

func TestPokerServer(t *testing.T) {

	testStorage := TestStorage{
		store: map[string]int{
			"dev":   10,
			"Rahul": 20,
		},
	}

	t.Run("test get for dev", func(t *testing.T) {

		res, req := MakeGetRequest("/players/dev")

		server := pokerhttpserver.PlayerServer{
			Storage: &testStorage,
		}
		server.ServeHTTP(res, req)

		got := res.Body.String()
		want := "10"

		AssertStrings(t, got, want)
		if res.Code != http.StatusOK {
			t.Error("Expected Error Status Not Found")
		}

	})

	t.Run("test get for rahul", func(t *testing.T) {
		res, req := MakeGetRequest("/players/Rahul")
		server := pokerhttpserver.PlayerServer{
			Storage: &testStorage,
		}
		server.ServeHTTP(res, req)
		got := res.Body.String()
		want := "20"
		AssertStrings(t, got, want)
		if res.Code != http.StatusOK {
			t.Error("Expected Error Status Not Found")
		}
	})

	t.Run("test get for invalid user", func(t *testing.T) {
		res, req := MakeGetRequest("/player/SomeUndefinedRandomShitUser")
		server := pokerhttpserver.PlayerServer{
			Storage: &testStorage,
		}
		server.ServeHTTP(res, req)
		if res.Code != http.StatusNotFound {
			t.Error("Expected Error Status Not Found")
		}
	})

	t.Run("Make a POST request for dev", func(t *testing.T) {
		res, req := MakePostRequest("/players/dev")
		server := pokerhttpserver.PlayerServer{
			Storage: &testStorage,
		}
		server.ServeHTTP(res, req)
		if res.Code != http.StatusAccepted {
			t.Errorf("%d != %d", res.Code, http.StatusAccepted)
		}
		if testStorage.Calls[0] != "dev" {
			t.Errorf("Not same got %v", testStorage.Calls)
		}
	})

}

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%q /= %q", got, want)
	}
}

func MakeGetRequest(path string) (*httptest.ResponseRecorder, *http.Request) {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	res := httptest.NewRecorder()
	return res, req
}

func MakePostRequest(path string) (*httptest.ResponseRecorder, *http.Request) {
	req, _ := http.NewRequest(http.MethodPost, path, nil)
	res := httptest.NewRecorder()
	return res, req
}
