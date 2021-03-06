package game

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	BoardHeight = 6
	BoardWidth  = 7
)

//constants to be used as Game.Status
const Broken = "you broke it!"
const waiting = "Waiting for 2nd player"
const gameprogress = "Game is in progress!"
const gameover = "Game has finished!"

//Game holds all the game data
type Game struct {
	Status             string                        `json:"status"`
	GameBoard          [BoardHeight][BoardWidth]Slot `json:"gameboard"`
	PlayerSymbols      []string                      `json:"playersymbols"`
	Heights            [BoardWidth]int
	NumMoves           int
	IsStarted          bool
	IsComplete         bool
	PlayersTurn        int `json:"playersturn"`
	NumPlayers         int
	LastMove           [2]int //row, col
	ComputerDifficulty int    `json:"computerdifficulty"`
}

//struct for each Slot on the board
type Slot struct {
	Active bool   `json:"active"`
	Symbol string `json:"symbol"`
}

//NewGame initializes new game with base values
func NewGame() *Game {
	game := Game{
		Status:             "empty status",
		GameBoard:          newGameBoard(),
		PlayerSymbols:      []string{0: "X", 1: "O"},
		Heights:            [BoardWidth]int{},
		NumMoves:           0,
		IsStarted:          false,
		IsComplete:         false,
		NumPlayers:         0,
		PlayersTurn:        0,
		ComputerDifficulty: 3,
	}
	return &game
}
func newGameBoard() [BoardHeight][BoardWidth]Slot {
	return [BoardHeight][BoardWidth]Slot{}
}

//AddPlayer adds a new player to the game
func (game *Game) AddPlayer() {
	game.NumPlayers++
	switch game.NumPlayers {
	case 1:
		game.Status = waiting
	case 2:
		game.Status = gameprogress
		game.IsStarted = true
	}
}

//StringBoard returns 6 string representation of board
func (game *Game) StringBoard() (string1, string2, string3, string4, string5, string6 string) {
	output := make([]string, 6)
	for i := 0; i < BoardHeight; i++ {
		for j := 0; j < BoardWidth; j++ {
			output[i] += ("| ")
			if game.GameBoard[i][j].Active == false {
				output[i] += ("__")
			} else {
				output[i] += game.GameBoard[i][j].Symbol
			}
			output[i] += (" |")
		}
		//output += "\n"
	}
	return output[0], output[1], output[2], output[3], output[4], output[5]

}

//FastPrint was used for debugging to see the board quickly
func FastPrint(string1, string2, string3, string4, string5, string6 string) {
	fmt.Println(string1)
	fmt.Println(string2)
	fmt.Println(string3)
	fmt.Println(string4)
	fmt.Println(string5)
	fmt.Println(string6)
}

//MakeMove checks if criteria is met and then updates Game
func (game *Game) MakeMove(playerNum int, move int) {
	if game.Heights == [7]int{6, 6, 6, 6, 6, 6, 6} {
		game.Status = gameover
		game.IsComplete = true
	}
	if game.isPlayersTurn(playerNum) {
		if game.IsValidMove(move) {

			height := game.Heights[move]

			game.GameBoard[5-height][move].Symbol = game.PlayerSymbols[playerNum]
			game.GameBoard[5-height][move].Active = true
			if game.CheckWinner(playerNum, move, 5-game.Heights[move]) {
				game.Status = gameover
				game.IsComplete = true
			}
			game.Heights[move]++
			game.LastMove = [2]int{game.Heights[move], move}

			game.switchPlayersTurn(playerNum)
			game.NumMoves++

		}
	}
}

func (game *Game) isPlayersTurn(playerNum int) bool {
	return playerNum == game.PlayersTurn
}

func (game *Game) IsValidMove(move int) bool {
	validmoves := []int{0, 1, 2, 3, 4, 5, 6}
	if game.Heights[move] == BoardHeight {
		return false
	}
	for _, v := range validmoves {
		if move == v {
			return true
		}
	}
	return false
}

func (game *Game) switchPlayersTurn(playerNum int) {
	switch playerNum {
	case 0:
		game.PlayersTurn = 1
	case 1:
		game.PlayersTurn = 0
	}
}

//CheckWinner checks if the last move played ended the game
func (game *Game) CheckWinner(player int, col int, row int) bool {
	symbol := game.PlayerSymbols[player]
	numToWin := 3

	//check horiz
	// check to the left
	for i := 1; col-i >= 0; i++ {
		if game.GameBoard[row][col-i].Symbol == symbol {
			numToWin--
		} else {
			break
		}
	}
	//check to the right
	for j := 1; col+j <= BoardWidth-1; j++ {

		if game.GameBoard[row][col+j].Symbol == symbol {
			numToWin--

		} else {
			break
		}
	}

	if numToWin <= 0 {
		return true
	}
	numToWin = 3
	//reset numToWin
	//check vertical
	//only need to check the 3 tokens below
	for i := 1; i <= 3; i++ {

		if row+i < 6 {
			if game.GameBoard[row+i][col].Symbol == symbol {
				numToWin--
			}
		} else {
			break
		}
	}
	if numToWin <= 0 {

		return true
	}
	numToWin = 3
	//reset numToWin
	//check diagonol
	//first check top to bottom left to right
	//check top left
	for i := 1; col-i >= 0 && row-i >= 0; i++ {
		if game.GameBoard[row-i][col-i].Symbol == symbol {
			numToWin--
		} else {
			break
		}
	}
	//check bottom right
	for j := 1; col+j <= BoardWidth-1 && row+j <= BoardHeight-1; j++ {
		if game.GameBoard[row+j][col+j].Symbol == symbol {
			numToWin--
		} else {
			break
		}
	}

	if numToWin <= 0 {

		return true
	}
	numToWin = 3
	//check top right direction
	for i := 1; col+i <= BoardWidth-1 && row-i >= 0; i++ {

		if game.GameBoard[row-i][col+i].Symbol == symbol {
			numToWin--
		} else {
			break
		}
	}

	//check bottom left direction
	for j := 1; col-j >= 0 && row+j <= BoardHeight-1; j++ {

		if game.GameBoard[row+j][col-j].Symbol == symbol {
			numToWin--
		} else {
			break
		}
	}
	if numToWin <= 0 {

		return true
	}
	return false
}

//JsonEncode turns Game struct into json
func (game *Game) JsonEncode() []byte {
	json, err := json.Marshal(game)
	if err != nil {
		log.Fatal("Error in marshalling json:", err)
	}

	return json
}
