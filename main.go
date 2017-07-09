package main

import (
	"log"
	"net/http"
)

func main() {
	// mygame := game.NewGame()
	//
	// var input int
	// for !mygame.IsComplete {
	// 	fmt.Println("input:")
	// 	fmt.Scanln(&input)
	// 	mygame.MakeMove(0, input)
	// 	fmt.Println(mygame.BoardAsString1)
	// 	fmt.Println(mygame.BoardAsString2)
	// 	fmt.Println(mygame.BoardAsString3)
	// 	fmt.Println(mygame.BoardAsString4)
	// 	fmt.Println(mygame.BoardAsString5)
	// 	fmt.Println(mygame.BoardAsString6)
	// 	mygame.MakeMove(1, minimax.Minimax(3, *mygame, 1))
	// 	fmt.Println(mygame.BoardAsString1)
	// 	fmt.Println(mygame.BoardAsString2)
	// 	fmt.Println(mygame.BoardAsString3)
	// 	fmt.Println(mygame.BoardAsString4)
	// 	fmt.Println(mygame.BoardAsString5)
	// 	fmt.Println(mygame.BoardAsString6)
	// }
	//fmt.Println("game over!")
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./c4-react/build/")))
	router.HandleFunc("/ws", handler)
	log.Printf("serving connect 4 live on localhost: 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
