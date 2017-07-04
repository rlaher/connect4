package main

import (
	"log"
	"net/http"
)

func main() {
	//preparing mux and server
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./c4-react/public/")))
	//router.Handle("/ws", wsHandler{})

	//serving
	log.Printf("serving connect 4 live on localhost: 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))

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
