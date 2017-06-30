package main

import "fmt"

const (
	BOARDHEIGHT = 6
	BOARDWIDTH  = 7
)

type Game struct {
	GameBoard  [BOARDHEIGHT][BOARDWIDTH]string //check these are in the right order
	heights    [BOARDWIDTH]int
	NumMoves   int
	IsComplete bool
	LastMove   []int //x pos, y pos
	LastPlayer string
}

// func (game *Game) MakeMove(col int) error {
// 	if game.GameBoard[col] { //check if column is full -> make a heights array
// 		return errors.New("Invalid Move: Column is full.")
// 	}
// 	game.GameBoard[col].append(true)
// 	return nil
// }

func (game *Game) PrintBoard() {

	// for a := BOARDHEIGHT - 1; a >= 0; a-- {
	// 	for i := 0; i < BOARDWIDTH; i++ {
	// 		fmt.Print("| ")
	// 		if game.GameBoard[a][i] == "" {
	// 			fmt.Print(" ")
	// 		} else {
	// 			fmt.Print(game.GameBoard[a][i])
	// 		}
	// 		fmt.Print(" |")
	//
	// 	}
	// 	fmt.Println()
	// }
	for i := 0; i < BOARDHEIGHT; i++ {
		for j := 0; j < BOARDWIDTH; j++ {
			fmt.Print("| ")
			if game.GameBoard[i][j] == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(game.GameBoard[i][j])
			}
			fmt.Print(" |")
		}
		fmt.Println()
	}
	//fmt.Print(game.GameBoard)
}
func (game *Game) CheckWinner() bool {
	// player := game.LastPlayer
	// row := game.LastMove[0]
	// col := game.LastMove[1]
	//
	// numToWin := 3

	//check horiz
	//check to the left
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

	mygame.PrintBoard()
	mygame.LastPlayer = "X"
	mygame.LastMove = []int{2, 0}
	a := mygame.CheckWinner()
	fmt.Print(a)
}
