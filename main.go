package main

import (
	// "fmt"
	pong "go-pong/game"
)

func main() {
	game := pong.NewGame()
	// fmt.Println(game)
	game.Run()
}
