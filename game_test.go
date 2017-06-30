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
		fmt.Print(myGame.PrintBoard())
		if actual != expected {
			t.Fatalf("Check Winner didn't work, expected %d, got %d", expected, actual)
		}
	}

}
