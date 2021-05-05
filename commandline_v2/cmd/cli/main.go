package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/mpfen/Learn-Go-with-Tests/commandline_v2"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	defer close()

	fmt.Println("Let' play poker")
	fmt.Println("Type {Name} wins to record a win")

	poker.NewCLI(store, os.Stdin).PlayPoker()
}
