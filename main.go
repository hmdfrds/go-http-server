package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
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
	request := string(buffer[:n])
	fmt.Println("Request received:\n", request)

	// Check for the requested method and path

	lines := strings.Split(request, "\r\n")
	if len(lines) > 0 {
		parts := strings.Fields(lines[0])
		if len(parts) > 1 {

			method, path := parts[0], parts[1]
			var response string
			if method == "GET" {
				if path == "/" {
					path = "/index.html"
				}
				response = serveFile(path)
			} else if method == "POST" {
				response = "HTTP/1.1 418 I'm a teapot\r\nContent-Type: text/plain\r\n\r\n418 - Will do later"
			}
			conn.Write([]byte(response))

		}
	}
}

func serveFile(path string) string {
	extension := strings.ToLower((filepath.Ext(path)))
	var contentType string
	switch extension {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	default:
		contentType = "text/plain"
	}

	buf, err := os.ReadFile("public" + path)
	if err != nil {
		return "HTTP/1.1 404 Not Founc\r\nContent-Type: text/plain\r\n\r\n404 - Not Found"
	}

	return fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: %s\r\n\r\n%s", contentType, string(buf))

}
