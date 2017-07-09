package minimax

import (
	"connect4/game"
	"fmt"
	"testing"
)

// func TestGetAvailableMoves(t *testing.T) {
// 	myGame := game.NewGame()
//
// 	for _, v := range testcasesGetAvailableMoves {
// 		myGame.Heights = v.heights
// 		expected := v.availmoves
// 		actual := GetAvailableMoves(*myGame)
//
// 		equality := true
//
// 		for i, v := range expected {
//
// 			if v != actual[i] {
//
// 				equality = false
// 				break
// 			}
// 		}
//
// 		if !equality {
// 			t.Fatalf("GetAvailMoves didn't work, expected %d, got %d", expected, actual)
// 		}
// 	}
//
// }
//
// func TestHorizStreaks(t *testing.T) {
// 	mygame := game.NewGame()
// 	for _, v := range testcasesStreaks {
// 		mygame.GameBoard = v.mygame
// 		for i := 2; i <= 4; i++ {
// 			actual := countHoriz(*mygame, v.playerNum, i)
// 			expected := v.horizstreaks[i]
// 			if actual != expected {
// 				t.Fatalf("Test Horiz didn't work, expected %d, got %d for streaklength %d", expected, actual, i)
// 			}
// 		}
// 		for j := 2; j <= 4; j++ {
// 			actual := countVert(*mygame, v.playerNum, j)
// 			expected := v.vertstreaks[j]
// 			if actual != expected {
// 				t.Fatalf("Test Vert didn't work, expected %d, got %d for streaklength %d", expected, actual, j)
// 			}
// 		}
// 		for k := 2; k <= 4; k++ {
// 			actual := countDiag(*mygame, v.playerNum, k)
// 			expected := v.diagstreaks[k]
// 			if actual != expected {
// 				t.Fatalf("Test Diag didn't work, expected %d, got %d for streaklength %d", expected, actual, k)
// 			}
// 		}
// 	}
// }

func TestSearch(t *testing.T) {
	mygame := game.NewGame()
	for _, v := range testcasesSearch {
		mygame.GameBoard = v.mygame
		mygame.Heights = v.heights
		a := BestMove(v.depth, *mygame, v.playerNum)
		fmt.Println(a)
	}
}
