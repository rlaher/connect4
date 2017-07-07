package minimax

import "connect4/game"

//Lets assume computer is always player 2, playing as "O"

//OKay so the root of the tree is our current game state
//the Children of this tree is all the possible game states that result from our move
//and the Children of those Children are all the possible game states that result from opp move

//Each node must contain: pointer to Parent, game state, pointer to Children (array)

type GameNode struct {
	Parent   *GameNode
	Self     *game.Game
	Children []*GameNode
	Score    float32
}

//PopulateChildren fills up the Children slice of all next potential moves
func (myGameNode *GameNode) PopulateChildren() {
	avail := GetAvailableMoves(myGameNode.Self)
	for a := range avail {
		tempgame := myGameNode.Self
		tempgame.MakeMove(myGameNode.Self.PlayersTurn, a)

		var tempGameNode GameNode
		tempGameNode.Parent = myGameNode
		tempGameNode.Self = tempgame
		tempGameNode.Score = float32(ScoreGame(tempgame)) //revisit this line
		myGameNode.Children = append(myGameNode.Children, &tempGameNode)
	}
	//now populate all the Children nodes
	for _, v := range myGameNode.Children {
		v.PopulateChildren()
	}
}

func (myGameNode *GameNode) CalcExpectedValue() float32 {
	avail := GetAvailableMoves(myGameNode.Self)
	var values []float32

	for i, _ := range avail {
		if myGameNode.Children[i].Score == 0 {
			myGameNode.Children[i].Score = myGameNode.Children[i].CalcExpectedValue()
		}
		values[i] = myGameNode.Children[i].Score
	}
	var output float32
	for _, v := range values {
		output = output + v
	}
	output = output / float32(len(values))
	return output
}

func GetAvailableMoves(game *game.Game) []int {
	var output []int
	for i := 0; i <= 6; i++ {
		if game.Heights[i] < 6 {
			output = append(output, i)
		}
	}
	return output
}
func ScoreGame(game *game.Game) int {
	//no winner case
	if !game.IsComplete {
		return 0
	}
	//it's p1's turn right now so computer had last move (and won)
	if game.PlayersTurn == 0 {
		return 10
	}
	return -10
}

func (myGameNode *GameNode) Minimax() int {
	myGameNode.PopulateChildren()
	avail := GetAvailableMoves(myGameNode.Self)
	var moveValues []float32

	for i, v := range myGameNode.Children {
		moveValues[i] = v.CalcExpectedValue()
	}
	var index int
	var max float32
	max = 0
	for i, v := range moveValues {
		if v >= max {
			index = i
		}
	}
	return avail[index]

}
