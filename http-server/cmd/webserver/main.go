package main

import (
	"log"
	"net/http"

	poker "github.com/gomesmf/go-learning/http-server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5500", server); err != nil {
		log.Fatalf("could not listen on port 5500 %v", err)
	}
}
