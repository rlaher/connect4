package minimax

import "connect4/game"

func GetAvailableMoves(game *game.Game) []int {
	var output []int
	for i := 0; i <= 6; i++ {
		if game.Heights[i] < 6 {
			output = append(output, i)
		}
	}
	return output
}
