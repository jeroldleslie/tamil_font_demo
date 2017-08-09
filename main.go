package main

import (
	"tamil_font_demo/server"
	"log"
)

func main() {
	// start the server
	serverErr := server.StartServer("8080")
	if serverErr != nil {
		log.Println("Error starting server!")
		log.Fatal(serverErr)
	}
}
