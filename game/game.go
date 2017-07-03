package game

const (
	BoardHeight = 6
	BoardWidth  = 7
)
const testing = "testing"
const waiting = "Waiting for 2nd player"
const gamestarts = "Game is starting!"
const gameover = "Game has finished!"

type Game struct {
	Status        string                        `json:"status"`
	GameBoard     [BoardHeight][BoardWidth]slot `json:"gameboard"`
	PlayerSymbols []string                      `json:"playersymbols"`

	Heights     [BoardWidth]int
	NumMoves    int
	IsStarted   bool
	IsComplete  bool
	playersTurn int
	LastMove    []int //col #, row #
	NumPlayers  int
	LastPlayer  string `json:"lastplayer"` //can probably delte
}

type slot struct {
	Active bool   `json:"active"`
	Symbol string `json:"symbol"`
}

func NewGame() *Game {
	game := Game{
		Status:        testing,
		GameBoard:     newGameBoard(),
		PlayerSymbols: []string{0: "X", 1: "O"},
		Heights:       [BoardWidth]int{},
		NumMoves:      0,
		IsStarted:     false,
		IsComplete:    false,
		LastMove:      []int{},
		LastPlayer:    "",
		NumPlayers:    0,
		playersTurn:   0,
	}
	return &game
}
func newGameBoard() [BoardHeight][BoardWidth]slot {
	return [BoardHeight][BoardWidth]slot{}
}

func (game *Game) addPlayer() {
	game.NumPlayers++
	switch game.NumPlayers {
	case 1:
		game.Status = waiting
	case 2:
		game.Status = gamestarts
		game.IsStarted = true
	}
}

func (game *Game) StringBoard() string {
	var output string
	for i := 0; i < BoardHeight; i++ {
		for j := 0; j < BoardWidth; j++ {
			output += ("| ")
			if game.GameBoard[i][j].Symbol == "" {
				output += (" ")
			} else {
				output += game.GameBoard[i][j].Symbol
			}
			output += (" |")
		}
		output += "\n"
	}
	return output

}

func (game *Game) MakeMove(playerNum int, move int) {
	if game.isPlayersTurn(playerNum) {
		if game.isValidMove(move) {

			height := game.Heights[move]

			game.GameBoard[5-height][move].Symbol = game.PlayerSymbols[playerNum]
			game.GameBoard[5-height][move].Active = true
			game.Heights[move]++
			game.switchPlayersTurn(playerNum)
			if game.CheckWinner() {
				game.Status = gameover
				game.IsComplete = true
			}
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

func (game *Game) CheckWinner() bool {
	player := game.LastPlayer
	col := game.LastMove[1]
	row := game.LastMove[0]
	numToWin := 3

	//check horiz
	// check to the left
	for i := 1; col-i >= 0; i++ {
		if game.GameBoard[row][col-i].Symbol == player {
			numToWin--
		} else {
			break
		}
	}
	//check to the right
	for j := 1; col+j <= BoardWidth-1; j++ {
		if game.GameBoard[row][col+j].Symbol == player {
			numToWin--
		} else {
			break
		}
	}

	if numToWin <= 0 {
		game.IsComplete = true
		return true
	}
	numToWin = 3
	//reset numToWin
	//check vertical
	//only need to check the 3 tokens below
	for i := 1; i <= 3; i++ {

		if row+i < 6 {
			if game.GameBoard[row+i][col].Symbol == player {
				numToWin--
			}
		} else {
			break
		}
	}
	if numToWin <= 0 {
		game.IsComplete = true

		return true
	}
	numToWin = 3
	//reset numToWin
	//check diagonol
	//first check top to bottom left to right
	//check top left
	for i := 1; col-i >= 0 && row-i >= 0; i++ {
		if game.GameBoard[row-i][col-i].Symbol == player {
			numToWin--
		} else {
			break
		}
	}
	//check bottom right
	for j := 1; col+j <= BoardWidth-1 && row+j <= BoardHeight-1; j++ {
		if game.GameBoard[row+j][col+j].Symbol == player {
			numToWin--
		} else {
			break
		}
	}

	if numToWin <= 0 {
		game.IsComplete = true

		return true
	}
	numToWin = 3
	//check top right direction
	for i := 1; col+i <= BoardWidth-1 && row-i >= 0; i++ {

		if game.GameBoard[row-i][col+i].Symbol == player {
			numToWin--
		} else {
			break
		}
	}

	//check bottom left direction
	for j := 1; col-j >= 0 && row+j <= BoardHeight-1; j++ {

		if game.GameBoard[row+j][col-j].Symbol == player {
			numToWin--
		} else {
			break
		}
	}
	if numToWin <= 0 {
		game.IsComplete = true

		return true
	}
	return false
}
