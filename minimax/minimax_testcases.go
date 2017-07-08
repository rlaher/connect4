package minimax

import "connect4/game"

var testcasesGetAvailableMoves = []struct {
	heights    [7]int
	availmoves []int
}{
	{
		[7]int{0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 2, 3, 4, 5, 6},
	},
	{
		[7]int{6, 6, 6, 6, 6, 6, 6},
		[]int{},
	},
	{
		[7]int{6, 5, 6, 6, 3, 6, 0},
		[]int{1, 4, 6},
	},
}

var testcasesStreaks = []struct {
	mygame       [game.BoardHeight][game.BoardWidth]game.Slot
	playerNum    int
	horizstreaks []int
	vertstreaks  []int
	diagstreaks  []int
}{
	{
		[game.BoardHeight][game.BoardWidth]game.Slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		0,
		[]int{0, 0, 1, 4, 0}, //offset by two for indexing ease
		[]int{0, 0, 1, 5, 0},
		[]int{0, 0, 8, 0, 2},
	},
}
