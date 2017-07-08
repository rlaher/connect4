package minimax

import (
	"connect4/game"
	"fmt"
	"math"
)

//depth is difficulty/how long it takes. higher = more difficult
func BestMove(depth int, mygame game.Game, playerNum int) int {
	potentialMoves := make(map[int]int) //maps column # of move to the "score"
	otherplayer := math.Abs(float64(1 - playerNum))

	for move := 0; move < 7; move++ {
		if mygame.IsValidMove(move) {
			temp := mygame
			temp.MakeMove(playerNum, move)
			potentialMoves[move] = search(depth-1, temp, int(otherplayer)) //returns value of other players optimal move
		}
	}
	alpha := -999999
	var output int
	for key, value := range potentialMoves {
		if value > alpha {
			alpha = value
			output = key
		}
	}
	fmt.Printf("The associated alpha value is %d", alpha)
	return output
}

//goes through tree and assigns values to nodes up to depth argument
func search(depth int, mygame game.Game, playerNum int) int {
	return 0
}

//stateValuation is assigns values to non end game board states
//this way we can minimax without recursing through every possible game state

func stateValuation(mygame game.Game, playerNum int) int {
	return 0
}

//countConsecutive returns how many streaks of streakLength exist
func countConsecutive(mygame game.Game, playerNum int, streakLength int) int {
	return 0
}

//countHoriz counts horiz streaks of streakLength
//only counts to the right to avoid double counting
//doesn't count streak of 3 as a streak of 2
func countHoriz(mygame game.Game, playerNum int, streakLength int) int {
	count := 0     //overall # of streaks of n length
	currCount := 0 //this needs to be exactly n so that we don't count streaks of 2 as streaks of 3
	symbol := mygame.PlayerSymbols[mygame.PlayersTurn]
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
							currCount = 0
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
	symbol := mygame.PlayerSymbols[mygame.PlayersTurn]
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
							currCount = 0
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
	symbol := mygame.PlayerSymbols[mygame.PlayersTurn]
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
							currCount = 0
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
							currCount = 0
						}
						currCount = 0 //reset currCount
					}

				}

			}
		}
	}

	return count
}

//Lets assume computer is always player 2, playing as "O"

//OKay so the root of the tree is our current game state
//the Children of this tree is all the possible game states that result from our move
//and the Children of those Children are all the possible game states that result from opp move

//Each node must contain: pointer to Parent, game state, pointer to Children (array)
//
// type GameNode struct {
// 	Parent   *GameNode
// 	Self     game.Game
// 	Children []*GameNode
// 	Score    float32
// }
//
// //PopulateChildren fills up the Children slice of all next potential moves
// func (myGameNode *GameNode) PopulateChildren() {
// 	avail := GetAvailableMoves(myGameNode.Self)
// 	for a := range avail {
//
// 		tempgame := myGameNode.Self
// 		tempgame.MakeMove(myGameNode.Self.PlayersTurn, a)
// 		//game.FastPrint(tempgame.BoardAsString1, tempgame.BoardAsString2, tempgame.BoardAsString3, tempgame.BoardAsString4, tempgame.BoardAsString5, tempgame.BoardAsString6)
//
// 		var tempGameNode GameNode
// 		tempGameNode.Parent = myGameNode
// 		tempGameNode.Self = tempgame
// 		tempGameNode.Score = ScoreGame(&tempgame)
// 		myGameNode.Children = append(myGameNode.Children, &tempGameNode)
// 	}
// }
//
// //SumNodes will evaluate all the children of one node
// //will assign them value based on their Children
// //the first child with the highest score will be returned
// //as the move to be played
// //kinda janky way to do this but lets see if it works
// func (myGameNode *GameNode) BestMove() int {
// 	avail := GetAvailableMoves(myGameNode.Self)
// 	sums := make([]float32, 10, 1000)
// 	var sum float32
// 	var que []*GameNode
// 	// que = append(que, myGameNode.Children...)
// 	for i, v := range myGameNode.Children {
// 		que = append(que, v)
// 		for len(que) > 0 {
// 			current := que[0] //take first item out of q, put its children in q
// 			que = que[1:]
// 			que = append(que, current.Children...)
// 			sum = sum + current.Score
// 		}
// 		sums[i] = sum
// 		sum = 0
// 	}
// 	temp := float32(0)
//
// 	var index int
// 	var flag bool
// 	//return the move that corresponds to highest sum
//
// 	for i, v := range sums {
// 		if v > temp {
// 			flag = true
// 			index = i
// 			temp = v
// 		}
// 	}
// 	if !flag { //no winning move in sight
// 		n := rand.Int() % len(avail)
// 		return avail[n]
// 	}
// 	return avail[index]
// }
//
// func (myGameNode *GameNode) CalcExpectedValue() float32 {
// 	avail := GetAvailableMoves(myGameNode.Self)
// 	var values []float32
//
// 	for i, _ := range avail {
// 		if myGameNode.Children[i].Score == 0 {
// 			fmt.Println("u make here")
// 			myGameNode.Children[i].Score = myGameNode.Children[i].CalcExpectedValue()
// 			fmt.Println("break here huh")
// 		}
// 		fmt.Println(i)
// 		fmt.Println(myGameNode.Children[i].Score)
// 		values[i] = myGameNode.Children[i].Score
// 		fmt.Println("where is this beaking")
// 	}
// 	var output float32
// 	for _, v := range values {
// 		output = output + v
// 	}
// 	output = output / float32(len(values))
// 	return output
// }
//
func GetAvailableMoves(game game.Game) []int {
	var output []int
	for i := 0; i <= 6; i++ {
		if game.Heights[i] < 6 {
			output = append(output, i)
		}
	}
	return output
}

// func ScoreGame(game *game.Game) float32 {
// 	//no winner case
// 	if !game.IsComplete {
// 		return 0
// 	}
// 	//it's p1's turn right now so computer had last move (and won)
// 	if game.PlayersTurn == 0 {
// 		return 10
// 	}
// 	return -10
// }
//
// func (myGameNode *GameNode) Minimax() int {
// 	myGameNode.PopulateChildren()
//
// 	//lets try getting 2 layers of moves
// 	for _, v := range myGameNode.Children {
// 		v.PopulateChildren()
// 	}
//
// 	return myGameNode.BestMove()
// }
