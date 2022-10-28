package main

import (
	"fara/internal/server"
	"log"
)

func main() {
	if err := server.Server(); err != nil {
		log.Println(err.Error())
		return
	}
}
