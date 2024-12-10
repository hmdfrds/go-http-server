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

func SearchHandler(r request.Request) response.Response {
	if r.Method != "GET" {
		return response.NewNotFoundResponse()
	}

	queryReqParts := strings.Split(r.Path, "?")
	if len(queryReqParts) < 2 {
		return response.NewResponse(200, "text/plain", "No Parameter Provided")
	}

	_, query := queryReqParts[0], queryReqParts[1]

	params := getParams(query)

	if len(params) <= 0 {
		return response.NewResponse(200, "text/plain", "Wrong parameters format")
	}

	if value, exist := params["q"]; exist {
		return response.NewResponse(200, "text/plain", "Searching for "+value)
	}

	return response.NewResponse(200, "text/plain", "parameter q not found")
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
		fmt.Printf("Error serving file: %s, %v\n", path, err)
		return response.NewInternalServerErrorResponse()
	}
	return response.NewResponse(http.StatusOK, getContentType(cleanPath), string(buf))
}

// Can easily call http.DetectContentType() to do this
func getContentType(path string) string {
	extension := strings.ToLower((filepath.Ext(path)))
	switch extension {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	default:
		return "application/octet-stream"
	}
}

func getParams(query string) map[string]string {
	queryParts := strings.Split(query, "&")

	if len(queryParts) < 1 {
		return nil
	}

	params := make(map[string]string)

	for _, query := range queryParts {
		kv := strings.SplitN(query, "=", 2)
		if len(kv) == 2 {
			params[kv[0]] = kv[1]
		}
	}
	return params
}
