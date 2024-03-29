package main

import (
	poker "learn-go-with-tests"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer close()

	game := poker.NewTexasHoldem(store, poker.BlindAlerterFunc(poker.Alerter))
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatal("problem creating player server", err)
	}

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatalf("could not listen on port 3000 %v", err)
	}
}
