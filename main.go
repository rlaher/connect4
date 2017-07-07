package main

import (
	"connect4/game"
	"connect4/minimax"
	"fmt"
)

func main() {
	mygame := game.NewGame()
	var myGameNode minimax.GameNode
	myGameNode.Self = *mygame
	var input int
	for !myGameNode.Self.IsComplete {
		fmt.Println("input:")
		fmt.Scanln(&input)
		myGameNode.Self.MakeMove(0, input)
		fmt.Println(myGameNode.Self.BoardAsString1)
		fmt.Println(myGameNode.Self.BoardAsString2)
		fmt.Println(myGameNode.Self.BoardAsString3)
		fmt.Println(myGameNode.Self.BoardAsString4)
		fmt.Println(myGameNode.Self.BoardAsString5)
		fmt.Println(myGameNode.Self.BoardAsString6)
		myGameNode.Self.MakeMove(1, myGameNode.Minimax())
		fmt.Println(myGameNode.Self.BoardAsString1)
		fmt.Println(myGameNode.Self.BoardAsString2)
		fmt.Println(myGameNode.Self.BoardAsString3)
		fmt.Println(myGameNode.Self.BoardAsString4)
		fmt.Println(myGameNode.Self.BoardAsString5)
		fmt.Println(myGameNode.Self.BoardAsString6)
	}

	// router := http.NewServeMux()
	// router.Handle("/", http.FileServer(http.Dir("./c4-react/build/")))
	// router.HandleFunc("/ws", handler)
	// log.Printf("serving connect 4 live on localhost: 8080")
	// log.Fatal(http.ListenAndServe("localhost:8080", router))
}
