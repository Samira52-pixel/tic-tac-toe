package main

import "tic-tac-toe/v1/tic-tac-toe/game"

func main() {
	for {
		if game.InitBoard() {
			break
		}
	}

	game.Play()
}
