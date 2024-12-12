package handler_test

import (
	"go-http-server/handler"
	"go-http-server/request"
	"os"
	"path/filepath"
	"testing"
)

func TestHelloHandler(t *testing.T) {

	// go out to the root folder
	// when run test, this directory will be the main
	dir, _ := os.Getwd()
	projectRoot := filepath.Join(dir, "..")
	os.Chdir(projectRoot)

	req := request.Request{Method: "GET", Path: "/hello/John"}
	res := handler.HelloHandler(req)

	if res.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}
	if res.ContentType != "text/html" {
		t.Errorf("Expected content type 'text/html', got '%s'", res.ContentType)
	}
}

func TestSearchHandler(t *testing.T) {
	req := request.Request{Method: "GET", Path: "/search?q=golang"}
	res := handler.SearchHandler(req)

	if res.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}

	if res.Body != "Searching for golang" {
		t.Errorf("Expected body 'Searching for golang', got %s", res.Body)
	}
}
