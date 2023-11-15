package main

import (
	"log"
	"net/http"

	psqlstorage "github.com/Rahul-NITD/SAGo/PokerServer/PSQLStorage"
	"github.com/Rahul-NITD/SAGo/PokerServer/pokerhttpserver"
)

func main() {

	dbstorage, _ := psqlstorage.CreateStorage(false)
	defer dbstorage.DisconnectDB()

	server := pokerhttpserver.NewPlayerServer(dbstorage)
	log.Fatal(http.ListenAndServe(":8000", http.Handler(server)))

}
