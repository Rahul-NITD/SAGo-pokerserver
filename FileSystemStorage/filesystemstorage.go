package filesystemstorage

import (
	"encoding/json"
	"io"

	"github.com/Rahul-NITD/SAGo/PokerServer/pokerhttpserver"
)

type FileSystemStorage struct {
  Databse io.Reader
}

func (fss *FileSystemStorage)GetLeague() []pokerhttpserver.Player {
  var league []pokerhttpserver.Player 
  json.NewDecoder(fss.Databse).Decode(&league)
  return league

}











































