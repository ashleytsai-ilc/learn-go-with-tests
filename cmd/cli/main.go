package main

import (
	"fmt"
	poker "learn-go-with-tests"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewTexasHoldem(store, poker.BlindAlerterFunc(poker.StdOutAlerter))
	cli := poker.NewCLI(os.Stdout, os.Stdin, game)

	cli.PlayPoker()
}
