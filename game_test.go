package main

import (
	"fmt"
	"testing"
)

var testcases = []struct {
	mygame [BOARDHEIGHT][BOARDWIDTH]string
	result bool
	pos    []int
	player string
}{
	{
		[BOARDHEIGHT][BOARDWIDTH]string{},
		false,
		[]int{0, 0},
		"O",
	},
	{
		[BOARDHEIGHT][BOARDWIDTH]string{
			{"O", "O", "O", "O"},
			{},
		},
		true,
		[]int{0, 0},
		"O",
	},
	{
		[BOARDHEIGHT][BOARDWIDTH]string{
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "X", "X", "O", "O", "O", "X"},
			{"X", "X", "X", "O", "O", "O", "X"},
			{"X", "X", "X", "O", "O", "O", "X"},
		},
		true,
		[]int{0, 0},
		"O",
	},
	{
		[BOARDHEIGHT][BOARDWIDTH]string{
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "O", "O", "X", "O", "X", "O"},
			{"O", "O", "O", "X", "X", "X", "O"},
			{"X", "X", "O", "O", "O", "O", "X"},
			{"X", "X", "X", "O", "X", "O", "X"},
			{"X", "X", "X", "O", "O", "O", "X"},
		},
		true,
		[]int{0, 0},
		"O",
	},
	{
		[BOARDHEIGHT][BOARDWIDTH]string{
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "O", "O", "X", "O", "X", "O"},
			{"O", "O", "O", "X", "X", "X", "O"},
			{"X", "X", "O", "O", "O", "O", "X"},
			{"X", "X", "X", "O", "X", "O", "X"},
			{"X", "X", "X", "O", "O", "O", "X"},
		},
		false,
		[]int{5, 4},
		"X",
	},
	{
		[BOARDHEIGHT][BOARDWIDTH]string{
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "O", "O", "X", "O", "X", "O"},
			{"O", "O", "X", "X", "X", "X", "O"},
			{"X", "X", "O", "X", "O", "O", "X"},
			{"X", "X", "X", "O", "X", "O", "X"},
			{"X", "X", "X", "O", "O", "O", "X"},
		},
		true,
		[]int{3, 3},
		"X",
	},
	{
		[BOARDHEIGHT][BOARDWIDTH]string{
			{"O", "O", "O", "O", "X", "X", "O"},
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "X", "X", "O", "O", "O", "X"},
			{"X", "X", "X", "O", "O", "O", "X"},
			{"X", "X", "X", "O", "O", "O", "X"},
		},
		true,
		[]int{0, 3},
		"O",
	},
	{
		[BOARDHEIGHT][BOARDWIDTH]string{
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "O", "O", "X", "X", "X", "O"},
			{"O", "X", "X", "O", "O", "O", "X"},
			{"X", "X", "X", "O", "O", "O", "X"},
			{"O", "X", "X", "O", "O", "O", "X"},
		},
		false,
		[]int{0, 6},
		"O",
	},
}

func TestCheckWin(t *testing.T) {
	var myGame Game
	for _, v := range testcases {
		myGame.GameBoard = v.mygame
		myGame.LastMove = v.pos
		myGame.LastPlayer = v.player

		expected := v.result
		actual := myGame.CheckWinner()
		fmt.Println("index:")
		fmt.Println(myGame.LastMove)
		fmt.Println("player symbol")
		fmt.Println(myGame.LastPlayer)
		fmt.Print(myGame.StringBoard())
		if actual != expected {
			t.Fatalf("Check Winner didn't work, expected %d, got %d", expected, actual)
		}
	}

}
