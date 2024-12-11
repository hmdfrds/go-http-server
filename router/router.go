package router

import (
	request "go-http-server/Request"
	response "go-http-server/Response"
	"strings"
)

type HandlerFunc func(request.Request) response.Response

type Router struct {
	handler map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{handler: make(map[string]HandlerFunc)}
}

func (r *Router) AddHandlerFunc(path string, handlerFunc HandlerFunc) {
	r.handler[path] = handlerFunc
}

func (r *Router) Handle(req request.Request) response.Response {
	
	for path, handlerFunc := range r.handler {

		if validRequest(path, req.Path) {
			return handlerFunc(req)
		}
	}
	return response.NewNotFoundResponse()
}

// Check if the request path is valid or not
// For now it can check for static, dynamic, and query path
func validRequest(route, reqPath string) bool {
	pathAndQuery := strings.SplitN(reqPath, "?", 2)
	reqPathOnly := pathAndQuery[0]

	routeParts := strings.Split(route, "/")
	reqParts := strings.Split(reqPathOnly, "/")

	if len(routeParts) != len(reqParts) {
		return false
	}

	for i, part := range routeParts {
		if part != reqParts[i] && !strings.HasPrefix(part, "{") {
			return false
		}
	}

	// Maybe I should allow every path to accept query
	if route == "/search" && len(pathAndQuery) > 1 {
		return true
	}

	return true
}
