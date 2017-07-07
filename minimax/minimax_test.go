package minimax

import (
	"connect4/game"
	"testing"
)

func TestGetAvailableMoves(t *testing.T) {
	myGame := game.NewGame()

	for _, v := range testcases {
		myGame.Heights = v.heights
		expected := v.availmoves
		actual := GetAvailableMoves(myGame)

		equality := true

		for i, v := range expected {

			if v != actual[i] {

				equality = false
				break
			}
		}

		if !equality {
			t.Fatalf("GetAvailMoves didn't work, expected %d, got %d", expected, actual)
		}
	}

}
