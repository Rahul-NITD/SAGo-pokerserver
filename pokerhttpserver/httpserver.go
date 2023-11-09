package pokerhttpserver

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerScoreStorage interface {
	GetPlayerScore(player string) (int, error)
	RecordWin(player string) error
}

type PlayerServer struct {
	Storage PlayerScoreStorage
}

func (p *PlayerServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	player := strings.TrimPrefix(req.URL.Path, "/players/")
	switch req.Method {
	case http.MethodPost:
		p.recordWin(res, player)
	case http.MethodGet:
		p.getScore(res, player)
	}
}

func (p *PlayerServer) getScore(res http.ResponseWriter, player string) {
	score, err := p.Storage.GetPlayerScore(player)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(res, score)
}

func (p *PlayerServer) recordWin(res http.ResponseWriter, player string) {
	err := p.Storage.RecordWin(player)
	if err != nil {
		res.WriteHeader(http.StatusNotModified)
		return
	}
	res.WriteHeader(http.StatusAccepted)
}
