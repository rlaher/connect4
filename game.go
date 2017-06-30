package main

import "fmt"

const (
	BOARDHEIGHT = 6
	BOARDWIDTH  = 7
)

type Game struct {
	GameBoard  [BOARDHEIGHT][BOARDWIDTH]string
	heights    [BOARDWIDTH]int
	NumMoves   int
	IsComplete bool
	LastMove   []int //col #, row #
	LastPlayer string
}

// func (game *Game) MakeMove(col int) error {
// 	if game.GameBoard[col] { //check if column is full -> make a heights array
// 		return errors.New("Invalid Move: Column is full.")
// 	}
// 	game.GameBoard[col].append(true)
// 	return nil
// }

func (game *Game) PrintBoard() string {
	var output string
	for i := 0; i < BOARDHEIGHT; i++ {
		for j := 0; j < BOARDWIDTH; j++ {
			output += ("| ")
			if game.GameBoard[i][j] == "" {
				output += (" ")
			} else {
				output += game.GameBoard[i][j]
			}
			output += (" |")
		}
		output += "\n"
	}
	return output

}

func (game *Game) CheckWinner() bool {
	player := game.LastPlayer
	col := game.LastMove[1]
	row := game.LastMove[0]

	numToWin := 3

	//check horiz
	// check to the left
	for i := 1; col-i >= 0; i++ {
		if game.GameBoard[row][col-i] == player {
			numToWin--
		} else {
			break
		}
	}
	//check to the right
	for j := 1; col+j <= BOARDWIDTH-1; j++ {
		if game.GameBoard[row][col+j] == player {
			numToWin--
		} else {
			break
		}
	}
	fmt.Print("num to win:")
	fmt.Println(numToWin)

	if numToWin <= 0 {
		return true
	} else {
		numToWin = 3
	} //reset numToWin

	//check vertical
	//only need to check the 3 tokens below
	for i := 1; i <= 3; i++ {
		if row+i < 7 {
			if game.GameBoard[row+i][col] == player {
				numToWin--
			}
		} else {
			break
		}
	}
	if numToWin <= 0 {
		return true
	} else {
		numToWin = 3
	} //reset numToWin
	//diagnol
	return false
}

func main() {
	var mygame Game
	mygame.GameBoard[0][0] = "X"
	mygame.GameBoard[0][1] = "O"
	mygame.GameBoard[0][2] = "O"
	mygame.GameBoard[0][3] = "O"
	mygame.GameBoard[0][4] = "O"

	board := mygame.PrintBoard()
	fmt.Print(board)
	mygame.LastPlayer = "O"
	mygame.LastMove = []int{0, 2}
	a := mygame.CheckWinner()
	fmt.Print(a)
}
