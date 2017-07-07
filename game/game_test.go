package game

import "testing"

func TestCheckWin(t *testing.T) {
	var myGame Game
	for _, v := range testcases {
		myGame.GameBoard = v.mygame
		myGame.PlayerSymbols = v.playersymbols
		expected := v.result

		actual := myGame.CheckWinner(v.player, v.pos[0], v.pos[1])

		if actual != expected {
			t.Fatalf("Check Winner didn't work, expected %d, got %d", expected, actual)
		}
	}

}
