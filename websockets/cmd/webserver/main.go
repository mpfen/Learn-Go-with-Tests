package main

import (
	"log"
	"net/http"

	poker "github.com/mpfen/Learn-Go-with-Tests/websockets"
)

const dbFileName = "game.db.json"

func main() {

	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}
	defer close()

	game := poker.NewGame(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
