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

	if route == "/search" && len(pathAndQuery) > 1 {
		return true
	}

	return true
}
