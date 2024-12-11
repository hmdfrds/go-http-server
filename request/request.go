package request

import (
	"fmt"
	"go-http-server/utils"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Body    string
}

func GetRequest(str string) (Request, error) {
	var body string

	requestParts := strings.Split(str, "\r\n\r\n")

	if len(requestParts) > 1 {
		body = utils.UnescapeString(requestParts[1])
	}

	requestLines := strings.Split(requestParts[0], "\r\n")
	if len(requestLines) < 1 {
		return Request{}, fmt.Errorf("invalid request: missing request line")
	}

	parts := strings.Split(requestLines[0], " ")
	if len(parts) < 3 {
		return Request{}, fmt.Errorf("invalid request line: %s", requestLines[0])
	}
	return Request{Method: parts[0], Path: parts[1], Version: parts[2], Body: body}, nil
}
