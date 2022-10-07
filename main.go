package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}
func main() {
	// fmt.Println("Hello")
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting Server on : 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
