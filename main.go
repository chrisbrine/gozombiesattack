package main

import (
	"go-zombies-attack/zombies"
)

func main() {
	game := zombies.NewGame()
	game.Run()
}