package filesystemstorage_test

import (
	"reflect"
	"strings"
	"testing"

	filesystemstorage "github.com/Rahul-NITD/SAGo/PokerServer/FileSystemStorage"
	"github.com/Rahul-NITD/SAGo/PokerServer/pokerhttpserver"
)

func TestFileSystemStorage(t *testing.T) {
	t.Run("Get league from reader", func(t *testing.T) {
		db := strings.NewReader(`[{name:"Rahul",wins:4},{name:"Akku",wins:3}]`)
		storage := filesystemstorage.FileSystemStorage{db}
		got := storage.GetLeague()
		want := []pokerhttpserver.Player{
			{
				Name: "Rahul",
				Wins: 4,
			},
			{
        Name: "Akku", 
        Wins: 3,
      },
		}
    if !reflect.DeepEqual(got, want) {
      t.Errorf("want %+v, got %+v", got, want)
    }
	})
}
