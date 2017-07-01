package player

import (
	"connect4/game"
	"errors"
	"fmt"
)

type Player struct {
	name   string
	marker string
	game   *game.Game
}

func (player *Player) MakeMove() error {
	var move int
	validmoves := []int{0, 1, 2, 3, 4, 5, 6}
	var validmove bool
	fmt.Printf("%s, What column would you like to play? (0-6)", player.name)
	fmt.Scan(&move)

	for _, v := range validmoves {
		if move == v {
			validmove = true
			break
		}
	}
	if !validmove {
		err := errors.New("Invalid Move! Try again")
		return err
	}
	if player.game.Heights[move] == game.BoardHeight {
		err := errors.New("Column is full! Try again")
		return err
	}
	height := player.game.Heights[move]

	player.game.GameBoard[5-height][move] = player.marker
	player.game.Heights[move] += 1
	player.game.LastPlayer = player.marker
	player.game.LastMove = []int{5 - height, move}
	fmt.Println("Move made!")
	fmt.Print(player.game.StringBoard())
	return nil

}
