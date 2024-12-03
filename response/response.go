package response

import (
	"fmt"
	"net/http"
)

type Response struct {
	StatusCode  int
	ContentType string
	Body        string
}

func NewResponse(statusCode int, contentType string, body string) Response {
	return Response{StatusCode: statusCode, ContentType: contentType, Body: body}
}
func NewNotFoundResponse() Response {
	return Response{StatusCode: 404, ContentType: "text/plain", Body: "404 - Not Found"}
}
func NewInternalServerErrorResponse() Response {
	return Response{StatusCode: 500, ContentType: "text/plain", Body: "500 - Internal Server Error"}
}

func NewBadRequestResponse() Response {
	return Response{StatusCode: 400, ContentType: "text/plain", Body: "400 - Bad Request"}
}

func (r *Response) String() string {
	return fmt.Sprintf("HTTP/1.1 %d %s\r\nContent-Type: %s\r\n\r\n %s", r.StatusCode, http.StatusText(r.StatusCode), r.ContentType, r.Body)
}
