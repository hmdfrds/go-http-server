package handler

import (
	"fmt"
	request "go-http-server/request"
	response "go-http-server/response"
	"go-http-server/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Just give them the HTML page if it's exist
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

// Just hello the input
// I don't have the page. This is the URL: http://localhost:8080/hello/{name}
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
		fmt.Println(err)
		return response.NewNotFoundResponse()
	}
	return response.NewResponse(http.StatusOK, getContentType("/hello.html"), fmt.Sprintf(string(buf), utils.UnescapeString(name)))

}

// Just return a plain page. Didn't make any search
func SearchHandler(r request.Request) response.Response {
	if r.Method != "GET" {
		return response.NewNotFoundResponse()
	}

	queryReqParts := strings.Split(r.Path, "?")
	if len(queryReqParts) < 2 {
		return response.NewOkTextPlainResponse("No Parameter Provided")
	}

	_, query := queryReqParts[0], queryReqParts[1]

	params := getParams(query)

	if len(params) <= 0 {
		return response.NewOkTextPlainResponse("Wrong parameters format")
	}

	if value, exist := params["q"]; exist {
		return response.NewOkTextPlainResponse("Searching for " + value)
	}

	return response.NewOkTextPlainResponse("parameter q not found")
}

// Return a result page base on the user input.
func SubmitFormHandler(r request.Request) response.Response {

	if r.Method != "POST" {
		return response.NewNotFoundResponse()
	}

	params := getParams(r.Body)

	if len(params) <= 0 {
		return response.NewOkTextPlainResponse("No parameters provided")
	}
	name := params["name"]
	email := params["email"]
	fmt.Println(r.Body)
	if name == "" || email == "" {
		return response.NewOkTextPlainResponse("Required parameters not found")
	}

	buf, err := os.ReadFile("public/result.html")
	if err != nil {
		return response.NewNotFoundResponse()
	}
	return response.NewResponse(http.StatusOK, getContentType("/result.html"), fmt.Sprintf(string(buf), name, email))
}

// Get the file from static folder
// Will cache it for 3 min
func StaticFileHandler(r request.Request) response.Response {
	if r.Method == "GET" && !strings.HasPrefix(r.Path, "/static") {
		return response.NewNotFoundResponse()
	}
	res := serveFile(r.Path)
	return res.SetCache(360)
}

// Get the file from public folder
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

// Separate the params in to map
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
