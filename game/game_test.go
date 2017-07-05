package game

import "testing"

var testcases = []struct {
	mygame        [BoardHeight][BoardWidth]slot
	playersymbols []string
	result        bool
	pos           []int
	player        int
}{
	{
		[BoardHeight][BoardWidth]slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "O"}},
		},
		[]string{0: "X", 1: "O"},
		true,
		[]int{0, 0},
		1,
	},

	{
		[BoardHeight][BoardWidth]slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},

		[]string{0: "X", 1: "O"},

		true,
		[]int{0, 0},
		1,
	},
	{
		[BoardHeight][BoardWidth]slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		[]string{0: "X", 1: "O"},
		true,
		[]int{0, 0},
		1,
	},
	{
		[BoardHeight][BoardWidth]slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		[]string{0: "X", 1: "O"},
		false,
		[]int{5, 4},
		0,
	},
	{
		[BoardHeight][BoardWidth]slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "X"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		[]string{0: "X", 1: "O"},
		true,
		[]int{3, 3},
		1,
	},
	{
		[BoardHeight][BoardWidth]slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		[]string{0: "X", 1: "O"},
		true,
		[]int{0, 3},
		1,
	},
	{
		[BoardHeight][BoardWidth]slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		[]string{0: "X", 1: "O"},
		true,
		[]int{6, 0},
		0,
	},
}

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
