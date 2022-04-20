package main

import (
	"log"
	"net/http"

	poker "github.com/gomesmf/go-learn-with-tests/http-server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server, err := poker.NewPlayerServer(store)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":5500", server); err != nil {
		log.Fatalf("could not listen on port 5500 %v", err)
	}
}
