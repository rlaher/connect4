package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// var mygame game.Game
	// player1 := player.Player{
	// 	Name:   "Bob",
	// 	Marker: "O",
	// 	Game:   &mygame,
	// }
	// player2 := player.Player{
	// 	Name:   "Joe",
	// 	Marker: "X",
	// 	Game:   &mygame,
	// }
	// for !mygame.IsComplete {
	// 	//player one gets to go
	// 	if !mygame.IsComplete {
	// 		err := player1.MakeMove()
	// 		for err != nil {
	// 			fmt.Println(err.Error())
	// 			err = player1.MakeMove()
	// 		}
	// 		mygame.IsComplete = mygame.CheckWinner()
	// 	}
	// 	//player2 gets to go
	// 	if !mygame.IsComplete {
	// 		err := player2.MakeMove()
	// 		for err != nil {
	// 			fmt.Println(err.Error())
	// 			err = player2.MakeMove()
	// 		}
	// 		mygame.IsComplete = mygame.CheckWinner()
	// 	}
	// }
	// fmt.Println("Game over!")
	gameRoom := NewGameRoom()
	go gameRoom.run()
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(gameRoom, w, r)
	})
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "home.html")

}