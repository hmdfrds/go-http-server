package handler

import (
	"fmt"
	request "go-http-server/Request"
	response "go-http-server/Response"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ServeHTML(r request.Request) response.Response {
	if r.Method != "GET" {
		return response.NewNotFoundResponse()
	}

	path := r.Path
	if r.Path == "/" {
		path = "/index"
	}

	return serveFile(path + ".html")

}

func HelloHandler(r request.Request) response.Response {
	if r.Method != "GET" {
		return response.NewNotFoundResponse()
	}

	parts := strings.Split(r.Path, "/")
	if len(parts) < 3 {
		return response.NewNotFoundResponse()
	}
	name := parts[2]

	buf, err := os.ReadFile("public/hello.html")
	if err != nil {
		return response.NewNotFoundResponse()
	}
	return response.NewResponse(http.StatusOK, getContentType("/hello.html"), fmt.Sprintf(string(buf), name))

}

func StaticFileHandler(r request.Request) response.Response {
	if r.Method == "GET" && !strings.HasPrefix(r.Path, "/static") {
		return response.NewNotFoundResponse()
	}
	res := serveFile(r.Path)
	return res.SetCache(360)
}

func serveFile(path string) response.Response {

	cleanPath := filepath.Clean("public" + path)
	buf, err := os.ReadFile(cleanPath)
	if err != nil {
		return handleError(err, path)
	}
	return response.NewResponse(http.StatusOK, getContentType(cleanPath), string(buf))
}

func getContentType(path string) string {
	extension := strings.ToLower((filepath.Ext(path)))
	// http.DetectContentType()
	switch extension {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	default:
		return "application/octet-stream"
	}
}

func handleError(err error, path string) response.Response {
	if err != nil {
		fmt.Printf("Error serving file: %s, %v\n", path, err)
		return response.NewInternalServerErrorResponse()
	}
	return response.NewNotFoundResponse()
}
