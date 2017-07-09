package minimax

import (
	"connect4/game"
	"fmt"
	"math"
)

//Minimax returns the col # of "best" move
//depth is how many moves ahead the AI looks
//proxy for difficulty/how long it takes.
// higher = more difficult
func Minimax(depth int, mygame game.Game, playerNum int) int {
	potentialMoves := make(map[int]int) //maps column # of move to the "score"
	otherplayer := int(math.Abs(float64(1 - playerNum)))

	for move := 0; move < 7; move++ {
		if mygame.IsValidMove(move) {
			temp := mygame
			temp.MakeMove(playerNum, move)
			//assign value to each move based on what other player is expected to do
			potentialMoves[move] = -traverse(depth-1, temp, otherplayer)
		}
	}
	score := -999999999999
	var output int
	//Maximise our score OR maximise our minimum losses
	//hence name minimax
	for key, value := range potentialMoves {
		if value > score {
			score = value
			output = key
		}
	}
	fmt.Printf("The associated score value is %d", score)
	fmt.Println()
	fmt.Print("choose col #")
	fmt.Println(output)
	return output
}

//traverse goes through tree and assigns values to nodes up to depth argument
func traverse(depth int, mygame game.Game, playerNum int) int {

	var potentialMoves []game.Game //maps column # of move to the "score"
	otherplayer := int(math.Abs(float64(1 - playerNum)))
	for move := 0; move < 7; move++ {
		if mygame.IsValidMove(move) {
			temp := mygame
			temp.MakeMove(playerNum, move)
			potentialMoves = append(potentialMoves, temp)
		}
	}
	//if below condition holds, we've reached a "leaf"
	//AKA game is over, or we've reached max depth
	if depth == 0 || mygame.IsComplete {

		return stateValuation(mygame, playerNum)
	}
	output := -9999999
	//recursively go through all possible potentialMoves/game scenarious
	//this keeps calling itself until the above condition ^^ is hit (leaf node)
	for _, v := range potentialMoves {
		//traverse is negative bc we are calling it on otherplayer
		//we want the other players statevaluation to be as close to zero as possible
		output = Max(output, -traverse(depth-1, v, otherplayer))
	}

	return output
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//stateValuation is assigns values to non end game board states
//this way we can minimax without recursing through every possible game state
//haven't studied the optimal weightings but lets make diagnols worth more
//because they're trickier
func stateValuation(mygame game.Game, playerNum int) int {
	otherplayer := int(math.Abs(float64(1 - playerNum)))
	//if other player can win, block it
	if countHoriz(mygame, otherplayer, 4)+countVert(mygame, otherplayer, 4)+countDiag(mygame, otherplayer, 4) > 0 {
		return -1000000
	}

	horizweight2 := 2
	vertweight2 := 2
	diagweight2 := 4
	horizweight3 := 8
	vertweight3 := 8
	diagweight3 := 16
	weight4 := 1000000

	value := 0
	value += horizweight2 * countHoriz(mygame, playerNum, 2)
	value += vertweight2 * countVert(mygame, playerNum, 2)
	value += diagweight2 * countDiag(mygame, playerNum, 2)
	value += horizweight3 * countHoriz(mygame, playerNum, 3)
	value += vertweight3 * countVert(mygame, playerNum, 3)
	value += diagweight3 * countDiag(mygame, playerNum, 3)
	value += weight4 * countHoriz(mygame, playerNum, 4)
	value += weight4 * countVert(mygame, playerNum, 4)
	value += weight4 * countDiag(mygame, playerNum, 4)
	return value

}

//countHoriz counts horiz streaks of streakLength
//only counts to the right to avoid double counting
//doesn't count streak of 3 as a streak of 2
func countHoriz(mygame game.Game, playerNum int, streakLength int) int {
	count := 0     //overall # of streaks of n length
	currCount := 0 //this needs to be exactly n so that we don't count streaks of 2 as streaks of 3
	symbol := mygame.PlayerSymbols[playerNum]
	//increment over every square
	for i := 0; i < game.BoardHeight; i++ { //rows
		for j := 0; j < game.BoardWidth; j++ { //columns
			if mygame.GameBoard[i][j].Active {
				//check that the symbol matches
				if mygame.GameBoard[i][j].Symbol == symbol {
					//make sure piece to left isn't the same symbol
					//so that we're not double counting
					if j == 0 || mygame.GameBoard[i][j-1].Symbol != symbol {
						currCount++
						for k := 1; k < 4; k++ {
							if j+k < game.BoardWidth { //make sure we don't go off board
								if mygame.GameBoard[i][j+k].Symbol == symbol {
									currCount++
								} else {
									break //break if different symbol
								}
							}
						}
						//only increment count if currentcount matches n
						if currCount == streakLength {
							count++
						}
						currCount = 0 //reset currCount
					}

				}

			}
		}
	}
	return count
}

//countVert will count downwards
func countVert(mygame game.Game, playerNum int, streakLength int) int {

	count := 0     //overall # of streaks of n length
	currCount := 0 //this needs to be exactly n so that we don't count streaks of 2 as streaks of 3
	symbol := mygame.PlayerSymbols[playerNum]

	//increment over every square
	for i := 0; i < game.BoardHeight; i++ { //rows
		for j := 0; j < game.BoardWidth; j++ { //columns

			if mygame.GameBoard[i][j].Active {
				//check that the symbol matches
				if mygame.GameBoard[i][j].Symbol == symbol {
					//make sure piece above isn't the same symbol
					//so that we're not double counting
					if i == 0 || mygame.GameBoard[i-1][j].Symbol != symbol {
						currCount++
						for k := 1; k < 4; k++ {
							if i+k < game.BoardHeight { //make sure we don't go off board
								if mygame.GameBoard[i+k][j].Symbol == symbol {
									currCount++
								} else {
									break //break if different symbol
								}
							}
						}
						//only increment count if currentcount matches n

						if currCount == streakLength {
							count++

						}
						currCount = 0 //reset currCount
					}

				}

			}
		}
	}

	return count
}

//countDiag will count downwards both to the right and the left
func countDiag(mygame game.Game, playerNum int, streakLength int) int {
	count := 0     //overall # of streaks of n length
	currCount := 0 //this needs to be exactly n so that we don't count streaks of 2 as streaks of 3
	symbol := mygame.PlayerSymbols[playerNum]
	//downwards to the right first
	//increment over every square
	for i := 0; i < game.BoardHeight; i++ { //rows
		for j := 0; j < game.BoardWidth; j++ { //columns
			if mygame.GameBoard[i][j].Active {
				//check that the symbol matches
				if mygame.GameBoard[i][j].Symbol == symbol {
					//make sure piece up and left to it isn't the same symbol
					//so that we're not double counting
					if i == 0 || j == 0 || mygame.GameBoard[i-1][j-1].Symbol != symbol {
						currCount++
						for k := 1; k < 4; k++ {
							if i+k < game.BoardHeight && j+k < game.BoardWidth { //make sure we don't go off board
								if mygame.GameBoard[i+k][j+k].Symbol == symbol {
									currCount++
								} else {
									break //break if different symbol
								}
							}
						}
						//only increment count if currentcount matches n
						if currCount == streakLength {
							count++
						}
						currCount = 0 //reset currCount
					}

				}

			}
		}
	}
	//now downwards to the left
	for i := 0; i < game.BoardHeight; i++ { //rows
		for j := 0; j < game.BoardWidth; j++ { //columns
			if mygame.GameBoard[i][j].Active {
				//check that the symbol matches
				if mygame.GameBoard[i][j].Symbol == symbol {
					//make sure piece up and left to it isn't the same symbol
					//so that we're not double counting
					if i == 0 || j == game.BoardWidth-1 || mygame.GameBoard[i-1][j+1].Symbol != symbol {
						currCount++
						for k := 1; k < 4; k++ {
							if i+k < game.BoardHeight && j-k >= 0 { //make sure we don't go off board
								if mygame.GameBoard[i+k][j-k].Symbol == symbol {
									currCount++
								} else {
									break //break if different symbol
								}
							}
						}
						//only increment count if currentcount matches n
						if currCount == streakLength {
							count++
						}
						currCount = 0 //reset currCount
					}

				}

			}
		}
	}
	return count
}

//GetAvailableMoves returns which columsn aren't full
func GetAvailableMoves(game game.Game) []int {
	var output []int
	for i := 0; i <= 6; i++ {
		if game.Heights[i] < 6 {
			output = append(output, i)
		}
	}
	return output
}
