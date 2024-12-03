package main

import (
	"fmt"
	request "go-http-server/Request"
	response "go-http-server/Response"
	"go-http-server/handlers"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	// Open the TCP socket
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("Listening on port 8080...")

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		// Handler connection
		go handleConnection(conn)
	}
}

func logRequest(req request.Request, res response.Response, duration time.Duration) {
	log.Printf("Method: %s, Path: %s, Status: %d(%s), Duration: %d ms", req.Method, req.Path, res.StatusCode, http.StatusText(res.StatusCode), duration.Milliseconds())
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read the incomming
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		log.Printf("Error reading: %v", err)
		return
	}

	// Parse the request
	req, err := request.GetRequest(string(buffer[:n]))
	if err != nil {
		res := response.NewBadRequestResponse()
		conn.Write([]byte(res.String()))
		logRequest(req, res, 0)
		return
	}

	// Check for the requested method and path
	start := time.Now()
	res := handlers.GetResponse(req)
	conn.Write([]byte(res.String()))

	logRequest(req, res, time.Since(start))
}
