package game

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	board  *Board
	player *Player
	reader *bufio.Reader
	state  GameState
}

func NewGame(board *Board, player *Player, reader *bufio.Reader) *Game {
	return &Game{
		board:  board,
		player: player,
		reader: reader,
		state:  playing,
	}
}

func (g *Game) updateState() {
	if g.board.CheckWin(g.player.Figure()) {
		if g.player.Figure() == cross {
			g.state = crossWin
		} else {
			g.state = noughtWin
		}

		return
	}

	if g.board.CheckDraw() {
		g.state = draw
	}
}

// Игровой цикл
func (g *Game) Play() {
	for g.state == playing {
		g.board.PrintBoard()

		fmt.Printf(
			"%s's turn. Enter row and column (e.g. 1 2) or q to quit: ",
			g.player.GetSymbol(),
		)

		input, err := g.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)

		if input == "q" {
			g.state = quit
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		row, err1 := strconv.Atoi(parts[0])
		col, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil ||
			row < 1 || col < 1 ||
			row > g.board.Size() || col > g.board.Size() {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		if g.board.SetSymbol(row-1, col-1, g.player.Figure()) {
			g.updateState()
			g.player.SwitchPlayer()
		} else {
			fmt.Println("This cell is already occupied!")
		}
	}

	g.board.PrintBoard()

	switch g.state {
	case crossWin:
		fmt.Println("X wins!")
	case noughtWin:
		fmt.Println("O wins!")
	case draw:
		fmt.Println("It's a draw!")
	default:
		fmt.Println("Game over!")
	}
}
