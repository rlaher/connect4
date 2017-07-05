package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./c4-react/build/")))
	router.HandleFunc("/ws", handler)
	log.Printf("serving connect 4 live on localhost: 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
