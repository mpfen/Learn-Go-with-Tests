package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/mpfen/Learn-Go-with-Tests/commandline"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play Play Poker")
	fmt.Println("Type {Name} to record a win")

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
