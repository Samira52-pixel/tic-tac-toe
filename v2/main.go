package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	tictactoe "tic-tac-toe/v2/game"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	boardSize := 0

	for {
		fmt.Print("Enter the size of the board (3-9): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		input = strings.TrimSpace(input)

		boardSize, err = strconv.Atoi(input)
		if err != nil {
			boardSize = tictactoe.BoardDefaultSize
		}

		if boardSize < tictactoe.BoardMinSize || boardSize > tictactoe.BoardMaxSize {
			fmt.Println("Invalid board size.")
			continue
		}

		break
	}

	board := tictactoe.NewBoard(boardSize)
	player := tictactoe.NewPlayer()
	game := tictactoe.NewGame(board, player, reader)

	game.Play()
}
