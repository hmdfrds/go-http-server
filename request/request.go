package request

import (
	"fmt"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
}

func GetRequest(str string) (Request, error) {
	requestLines := strings.Split(str, "\r\n")
	if len(requestLines) < 1 {
		return Request{}, fmt.Errorf("invalid request: missing request line")
	}

	parts := strings.Split(requestLines[0], " ")
	if len(parts) < 3 {
		return Request{}, fmt.Errorf("invalid request line: %s", requestLines[0])
	}
	return Request{Method: parts[0], Path: parts[1], Version: parts[2]}, nil
}
