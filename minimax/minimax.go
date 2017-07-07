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
			potentialMoves[move] = -search(depth-1, temp, int(otherplayer))
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
	fmt.Sprint("The associated alpha value is%d", alpha)
	return output

}

//goes through tree and assigns values to nodes up to depth argument
func search(depth int, mygame game.Game, playerNum int) int {
	return 0
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
