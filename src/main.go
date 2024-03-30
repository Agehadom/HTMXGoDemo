package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed index.html
var page []byte

func main() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	fs := http.FileServer(http.Dir("/src/"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	handle := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Pong with CSS")
		w.Write([]byte(page))
	}

	http.HandleFunc("/", handle)

	log.Printf("Server on port %v", s.Addr)
	log.Fatalln(http.ListenAndServe(s.Addr, nil))
}
