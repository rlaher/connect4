package game

import (
	"encoding/json"
	"log"
)

const (
	BoardHeight = 6
	BoardWidth  = 7
)
const Broken = "you broke it!"
const waiting = "Waiting for 2nd player"
const gameprogress = "Game is in progress!"
const gameover = "Game has finished!"

type Game struct {
	Status        string                        `json:"status"`
	GameBoard     [BoardHeight][BoardWidth]slot `json:"gameboard"`
	PlayerSymbols []string                      `json:"playersymbols"`
	Heights       [BoardWidth]int
	NumMoves      int
	IsStarted     bool
	IsComplete    bool
	playersTurn   int
	NumPlayers    int

	BoardAsString1 string `json:"boardasstring1"`
	BoardAsString2 string `json:"boardasstring2"`
	BoardAsString3 string `json:"boardasstring3"`
	BoardAsString4 string `json:"boardasstring4"`
	BoardAsString5 string `json:"boardasstring5"`
	BoardAsString6 string `json:"boardasstring6"`
}

type slot struct {
	Active bool   `json:"active"`
	Symbol string `json:"symbol"`
}

func NewGame() *Game {
	game := Game{
		Status:        "empty status",
		GameBoard:     newGameBoard(),
		PlayerSymbols: []string{0: "X", 1: "O"},
		Heights:       [BoardWidth]int{},
		NumMoves:      0,
		IsStarted:     false,
		IsComplete:    false,

		NumPlayers:     0,
		playersTurn:    0,
		BoardAsString1: "No board yet",
	}
	return &game
}
func newGameBoard() [BoardHeight][BoardWidth]slot {
	return [BoardHeight][BoardWidth]slot{}
}

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

func (game *Game) MakeMove(playerNum int, move int) {
	if game.isPlayersTurn(playerNum) {
		if game.isValidMove(move) {

			height := game.Heights[move]

			game.GameBoard[5-height][move].Symbol = game.PlayerSymbols[playerNum]
			game.GameBoard[5-height][move].Active = true
			if game.CheckWinner(playerNum, move, 5-game.Heights[move]) {
				game.Status = gameover
				game.IsComplete = true
			}
			game.Heights[move]++

			game.switchPlayersTurn(playerNum)
			game.NumMoves++
			game.BoardAsString1, game.BoardAsString2, game.BoardAsString3, game.BoardAsString4, game.BoardAsString5, game.BoardAsString6 = game.StringBoard()

		}
	}
}
func (game *Game) isPlayersTurn(playerNum int) bool {
	return playerNum == game.playersTurn
}

func (game *Game) isValidMove(move int) bool {
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

// switchPlayersTurn switches the playersTurn variable to the id of the other player
func (game *Game) switchPlayersTurn(playerNum int) {
	switch playerNum {
	case 0:
		game.playersTurn = 1
	case 1:
		game.playersTurn = 0
	}
}

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

func (game *Game) JsonEncode() []byte {
	json, err := json.Marshal(game)
	if err != nil {
		log.Fatal("Error in marshalling json:", err)
	}

	return json
}
