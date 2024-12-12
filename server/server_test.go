package server_test

import (
	"bytes"
	"fmt"
	request "go-http-server/request"
	response "go-http-server/response"
	"go-http-server/router"
	"go-http-server/server"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestServe(t *testing.T) {
	r := router.NewRouter()
	r.AddHandlerFunc("/", func(req request.Request) response.Response {
		return response.NewOkTextPlainResponse("Hello, World!")
	})

	s := &server.Server{Router: r}

	// Simulate a TCP connection
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	go s.Serve(server)

	client.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	buffer := make([]byte, 1024)
	n, _ := client.Read(buffer)

	expected := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello, World!"
	if !bytes.Contains(buffer[:n], []byte(expected)) {
		t.Errorf("Expected response '%s', got '%s'", expected, buffer[:n])
	}
}

func TestStartServer(t *testing.T) {

	// go out to the root folder
	// when run test, this directory will be the main
	dir, _ := os.Getwd()
	projectRoot := filepath.Join(dir, "..")
	os.Chdir(projectRoot)

	go server.StartServer()

	resp, err := http.Get("http://localhost:8080/hello/World")

	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	body, _ := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	buf, err := os.ReadFile("public/hello.html")
	if err != nil {
		t.Fatal("Error while reading hello.html\n", err)
	}

	expectedBody := fmt.Sprintf(string(buf), "World")

	if string(body) != expectedBody {
		t.Error("Expected response body is wrong")
	}
}
