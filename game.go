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
	// player := game.LastPlayer
	// row := game.LastMove[0]
	// col := game.LastMove[1]
	//
	// numToWin := 3

	// check horiz
	// // check to the left
	// for i := 1; xpos-i >= 0; i++ {
	// 	if game.GameBoard[xpos-i][ypos] == player {
	// 		numToWin--
	// 	} else {
	// 		break
	// 	}
	// }
	// //check to the right
	// for j := 1; xpos+j <= BOARDWIDTH-1; j++ {
	// 	if game.GameBoard[xpos+j][ypos] == player {
	// 		numToWin--
	// 	} else {
	// 		break
	// 	}
	// }
	// if numToWin <= 0 {
	// 	return true
	// }
	return false
	//check vertical

	//diagnol
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
	mygame.LastPlayer = "X"
	mygame.LastMove = []int{2, 0}
	a := mygame.CheckWinner()
	fmt.Print(a)
}
