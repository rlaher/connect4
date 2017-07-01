package game

const (
	BoardHeight = 6
	BoardWidth  = 7
)

type Game struct {
	GameBoard  [BoardHeight][BoardWidth]string
	Heights    [BoardWidth]int
	NumMoves   int
	IsComplete bool
	LastMove   []int //col #, row #
	LastPlayer string
}

func (game *Game) StringBoard() string {
	var output string
	for i := 0; i < BoardHeight; i++ {
		for j := 0; j < BoardWidth; j++ {
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
	for j := 1; col+j <= BoardWidth-1; j++ {
		if game.GameBoard[row][col+j] == player {
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
			if game.GameBoard[row+i][col] == player {
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
		if game.GameBoard[row-i][col-i] == player {
			numToWin--
		} else {
			break
		}
	}
	//check bottom right
	for j := 1; col+j <= BoardWidth-1 && row+j <= BoardHeight-1; j++ {
		if game.GameBoard[row+j][col+j] == player {
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

		if game.GameBoard[row-i][col+i] == player {
			numToWin--
		} else {
			break
		}
	}

	//check bottom left direction
	for j := 1; col-j >= 0 && row+j <= BoardHeight-1; j++ {

		if game.GameBoard[row+j][col-j] == player {
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
