package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameState int
type BoardField int

// Фигуры в клетке поля.
const (
	empty BoardField = iota
	cross
	nought
)

// Состояние игрового процесса.
const (
	playing GameState = iota
	draw
	crossWin
	noughtWin
	quit
)

var (
	board         [][]BoardField
	boardSize     int        = 3
	currentPlayer BoardField = cross
	state         GameState  = playing
	reader                   = bufio.NewReader(os.Stdin)
)

// InitBoard создаёт игровое поле размером от 3 × 3 до 9 × 9.
func InitBoard() bool {
	fmt.Print("Enter the size of the board (3-9): ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return false
	}

	input = strings.TrimSpace(input)

	size, err := strconv.Atoi(input)
	if err != nil {
		// Если введено не число, используется предыдущий размер.
		size = boardSize
	}

	if size < 3 || size > 9 {
		fmt.Println("Invalid board size.")
		return false
	}

	boardSize = size
	board = make([][]BoardField, boardSize)

	for i := range board {
		board[i] = make([]BoardField, boardSize)
	}

	currentPlayer = cross
	state = playing

	return true
}

// PrintBoard выводит текущее состояние игрового поля.
func PrintBoard() {
	fmt.Print("  ")

	for i := range boardSize {
		fmt.Printf("%d ", i+1)
	}

	fmt.Println()

	for i := range boardSize {
		fmt.Printf("%d ", i+1)

		for j := range boardSize {
			switch board[i][j] {
			case empty:
				fmt.Print(". ")
			case cross:
				fmt.Print("X ")
			case nought:
				fmt.Print("O ")
			}
		}

		fmt.Println()
	}
}

// makeMove проверяет и выполняет ход.
func makeMove(x, y int) bool {
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		return false
	}

	if board[x][y] != empty {
		return false
	}

	board[x][y] = currentPlayer
	return true
}

// checkWin проверяет, выиграл ли указанный игрок.
func checkWin(player BoardField) bool {
	// Проверка строк и столбцов.
	for i := range boardSize {
		rowWin := true
		colWin := true

		for j := range boardSize {
			if board[i][j] != player {
				rowWin = false
			}

			if board[j][i] != player {
				colWin = false
			}
		}

		if rowWin || colWin {
			return true
		}
	}

	// Проверка главной диагонали.
	mainDiag := true

	for i := range boardSize {
		if board[i][i] != player {
			mainDiag = false
			break
		}
	}

	if mainDiag {
		return true
	}

	// Проверка побочной диагонали.
	antiDiag := true

	for i := range boardSize {
		if board[i][boardSize-i-1] != player {
			antiDiag = false
			break
		}
	}

	return antiDiag
}

// checkDraw проверяет, заполнено ли всё поле.
func checkDraw() bool {
	for i := range boardSize {
		for j := range boardSize {
			if board[i][j] == empty {
				return false
			}
		}
	}

	return true
}

// switchPlayer передаёт ход следующему игроку.
func switchPlayer() {
	if currentPlayer == cross {
		currentPlayer = nought
	} else {
		currentPlayer = cross
	}
}

// updateState обновляет состояние игры.
func updateState() {
	if checkWin(currentPlayer) {
		if currentPlayer == cross {
			state = crossWin
		} else {
			state = noughtWin
		}
	} else if checkDraw() {
		state = draw
	}
}

// Play запускает игровой цикл.
func Play() {
	PrintBoard()

	for state == playing {
		playerSymbol := "X"

		if currentPlayer == nought {
			playerSymbol = "O"
		}

		fmt.Printf(
			"%s's turn. Enter row and column (e.g. 1 2): ",
			playerSymbol,
		)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)

		if input == "q" {
			state = quit
			break
		}

		parts := strings.Fields(input)

		if len(parts) != 2 {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		row, err1 := strconv.Atoi(parts[0])
		col, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil || !makeMove(row-1, col-1) {
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		updateState()
		PrintBoard()

		switch state {
		case crossWin:
			fmt.Println("X wins!")
		case noughtWin:
			fmt.Println("O wins!")
		case draw:
			fmt.Println("It's a draw!")
		default:
			switchPlayer()
		}
	}
}
