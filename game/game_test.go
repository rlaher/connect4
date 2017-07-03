package game

//
// import (
// 	"fmt"
// 	"testing"
// )
//
// //FIX LATER
// var testcases = []struct {
// 	mygame [BoardHeight][BoardWidth]string
// 	result bool
// 	pos    []int
// 	player string
// }{
// 	{
// 		[BoardHeight][BoardWidth]string{},
// 		false,
// 		[]int{0, 0},
// 		"O",
// 	},
// 	{
// 		[BoardHeight][BoardWidth]string{
// 			{"O", "O", "O", "O"},
// 			{},
// 		},
// 		true,
// 		[]int{0, 0},
// 		"O",
// 	},
// 	{
// 		[BoardHeight][BoardWidth]string{
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "X", "X", "O", "O", "O", "X"},
// 			{"X", "X", "X", "O", "O", "O", "X"},
// 			{"X", "X", "X", "O", "O", "O", "X"},
// 		},
// 		true,
// 		[]int{0, 0},
// 		"O",
// 	},
// 	{
// 		[BoardHeight][BoardWidth]string{
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "O", "O", "X", "O", "X", "O"},
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"X", "X", "O", "O", "O", "O", "X"},
// 			{"X", "X", "X", "O", "X", "O", "X"},
// 			{"X", "X", "X", "O", "O", "O", "X"},
// 		},
// 		true,
// 		[]int{0, 0},
// 		"O",
// 	},
// 	{
// 		[BoardHeight][BoardWidth]string{
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "O", "O", "X", "O", "X", "O"},
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"X", "X", "O", "O", "O", "O", "X"},
// 			{"X", "X", "X", "O", "X", "O", "X"},
// 			{"X", "X", "X", "O", "O", "O", "X"},
// 		},
// 		false,
// 		[]int{5, 4},
// 		"X",
// 	},
// 	{
// 		[BoardHeight][BoardWidth]string{
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "O", "O", "X", "O", "X", "O"},
// 			{"O", "O", "X", "X", "X", "X", "O"},
// 			{"X", "X", "O", "X", "O", "O", "X"},
// 			{"X", "X", "X", "O", "X", "O", "X"},
// 			{"X", "X", "X", "O", "O", "O", "X"},
// 		},
// 		true,
// 		[]int{3, 3},
// 		"X",
// 	},
// 	{
// 		[BoardHeight][BoardWidth]string{
// 			{"O", "O", "O", "O", "X", "X", "O"},
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "X", "X", "O", "O", "O", "X"},
// 			{"X", "X", "X", "O", "O", "O", "X"},
// 			{"X", "X", "X", "O", "O", "O", "X"},
// 		},
// 		true,
// 		[]int{0, 3},
// 		"O",
// 	},
// 	{
// 		[BoardHeight][BoardWidth]string{
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "O", "O", "X", "X", "X", "O"},
// 			{"O", "X", "X", "O", "O", "O", "X"},
// 			{"X", "X", "X", "O", "O", "O", "X"},
// 			{"O", "X", "X", "O", "O", "O", "X"},
// 		},
// 		false,
// 		[]int{0, 6},
// 		"O",
// 	},
// }
//
// func TestCheckWin(t *testing.T) {
// 	var myGame Game
// 	for _, v := range testcases {
// 		myGame.GameBoard = v.mygame
// 		myGame.LastMove = v.pos
// 		myGame.LastPlayer = v.player
//
// 		expected := v.result
// 		actual := myGame.CheckWinner()
// 		fmt.Println("index:")
// 		fmt.Println(myGame.LastMove)
// 		fmt.Println("player symbol")
// 		fmt.Println(myGame.LastPlayer)
// 		fmt.Print(myGame.StringBoard())
// 		if actual != expected {
// 			t.Fatalf("Check Winner didn't work, expected %d, got %d", expected, actual)
// 		}
// 	}
//
// }
