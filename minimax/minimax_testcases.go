package minimax

var testcases = []struct {
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
