package player

import (
	"connect4/game"
	"errors"
	"fmt"
)

type Player struct {
	Name   string
	Marker string
	Game   *game.Game
}

func (player *Player) MakeMove(move int) error {
	//var move int
	validmoves := []int{0, 1, 2, 3, 4, 5, 6}
	var validmove bool

	// fmt.Printf("%s, What column would you like to play? (0-6)", player.Name)
	// fmt.Scan(&move)

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
	if player.Game.Heights[move] == game.BoardHeight {
		err := errors.New("Column is full! Try again")
		return err
	}
	height := player.Game.Heights[move]

	player.Game.GameBoard[5-height][move] = player.Marker
	player.Game.Heights[move]++
	player.Game.LastPlayer = player.Marker
	player.Game.LastMove = []int{5 - height, move}
	fmt.Println("Move made!")
	fmt.Print(player.Game.StringBoard())
	return nil

}
func (player *Player) BotMove() {
	//bot should randomly choose column
	player.MakeMove(0)
}
