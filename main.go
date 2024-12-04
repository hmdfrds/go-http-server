package main

import (
	"fmt"
	"go-http-server/handler"
	"go-http-server/router"
	"go-http-server/server"
	"log"
	"net"
)

func main() {

	// Create router
	router := router.NewRouter()
	router.AddHandlerFunc("/", handler.ServeHTML)
	router.AddHandlerFunc("/home", handler.ServeHTML)
	router.AddHandlerFunc("/about", handler.ServeHTML)
	router.AddHandlerFunc("/hello/{name}", handler.HelloHandler)
	router.AddHandlerFunc("/static/style.css", handler.StaticFileHandler)
	// Create server
	server := &server.Server{Router: router}

	// Open the TCP socket
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("Listening on port 8080...")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		// Handler connection
		go server.Serve(conn)
	}
}
