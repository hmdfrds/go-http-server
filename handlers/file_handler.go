package handlers

import (
	request "go-http-server/Request"
	response "go-http-server/Response"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetResponse(r request.Request) response.Response {

	if r.Method == "GET" {
		if r.Path == "/" {
			r.Path = "/index.html"
		}
		return serveFile(r.Path)
	}

	return response.NewNotFoundResponse()

}

func serveFile(path string) response.Response {
	extension := strings.ToLower((filepath.Ext(path)))

	// http.DetectContentType()

	var contentType string
	switch extension {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	default:
		contentType = "application/octet-stream"
	}

	buf, err := os.ReadFile("public" + path)
	if err != nil {
		return response.NewNotFoundResponse()
	}

	return response.NewResponse(http.StatusOK, contentType, string(buf))

}
