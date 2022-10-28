package server

import (
	"fara/internal/handlers"
	"fmt"
	"net/http"
)

func Server() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.MainPage)
	mux.HandleFunc("/send", handlers.PostSend)
	mux.HandleFunc("/delay-send", handlers.PostDelay)
	mux.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./template/"))))
	fmt.Println("Server listening...")
	fmt.Println("http://localhost:7777")
	return http.ListenAndServe(":7777", mux)
}
