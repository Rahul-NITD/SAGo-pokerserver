package psqlstorage_test

import (
	"reflect"
	"testing"

	psqlstorage "github.com/Rahul-NITD/SAGo/PokerServer/PSQLStorage"
	"github.com/Rahul-NITD/SAGo/PokerServer/pokerhttpserver"
)

func TestPSQLStorage(t *testing.T) {

	db, err := psqlstorage.Connect()
	if err != nil {
		t.Fatalf("Could not connect to DB, %q", err.Error())
	}
	defer db.Close()

	t.Run("test creating table", func(t *testing.T) {

		err := psqlstorage.CreateTable(db, true)

		if err != nil {
			t.Fatalf("Error creating table, %q", err)
		}

	})

	t.Run("Test inserting into table", func(t *testing.T) {

		player := pokerhttpserver.Player{
			Name: "Rahul",
			Wins: 4,
		}

		err := psqlstorage.InsertPlayer(db, player)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Read League", func(t *testing.T) {

		psqlstorage.InsertPlayer(db, pokerhttpserver.Player{
			Name: "Akku",
			Wins: 5,
		})
		want := []pokerhttpserver.Player{
			{
				Name: "Akku",
				Wins: 5,
			},
			{
				Name: "Rahul",
				Wins: 4,
			},
		}

		got, err := psqlstorage.ReadLeague(db)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Got %v != want %v", got, want)
		}

		if err != nil {
			t.Fatalf("Error occured, %q", err.Error())
		}

	})

	t.Run("Test Update values", func(t *testing.T) {
		psqlstorage.UpdateValues(db, "Akku")
		want := []pokerhttpserver.Player{
			{
				Name: "Akku",
				Wins: 6,
			},
			{
				Name: "Rahul",
				Wins: 4,
			},
		}

		got, err := psqlstorage.ReadLeague(db)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

		if err != nil {
			t.Error(err.Error())
		}

	})

	t.Run("Test get score", func(t *testing.T) {
		score, err := psqlstorage.GetScore(db, "Rahul")
		if err != nil {
			t.Fatal(err)
		}
		if score != 4 {
			t.Errorf("Got %d, want %d", score, 4)
		}
	})

}

func TestIntegration(t *testing.T) {

	storage, err := psqlstorage.CreateStorage(true)
	if err != nil {
		t.Errorf("Error occured, %q", err.Error())
	}
	defer storage.DisconnectDB()

	t.Run("Test Interface", func(t *testing.T) {

		_, err := storage.GetPlayerScore("Rahul")

		if err.Error() != "Player Does not Exist" {
			t.Fatalf("This is a new error, err : %q", err.Error())
		}

	})

	t.Run("Test Record Win", func(t *testing.T) {
		err := storage.RecordWin("Rahul")
		if err != nil {
			t.Fatal(err)
		}
		score, err := storage.GetPlayerScore("Rahul")
		if err != nil {
			t.Fatal(err)
		}
		if score != 1 {
			t.Errorf("got %d, want %d", score, 1)
		}
	})

	for i := 0; i < 6; i++ {
		storage.RecordWin("Akku")
	}

	t.Run("Test League", func(t *testing.T) {
		got, err := storage.GetLeague()
		if err != nil {
			t.Fatal(err)
		}
		want := []pokerhttpserver.Player{
			{
				Name: "Akku",
				Wins: 6,
			},
			{
				Name: "Rahul",
				Wins: 1,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
