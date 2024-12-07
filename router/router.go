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

		if isDynamicMatch(path, req.Path) {
			return handlerFunc(req)
		}
	}
	return response.NewNotFoundResponse()
}

func isDynamicMatch(route, reqPath string) bool {
	routeParts := strings.Split(route, "/")
	reqParts := strings.Split(reqPath, "/")

	if len(routeParts) != len(reqParts) {
		return false
	}

	for i, part := range routeParts {
		if part != reqParts[i] && !strings.HasPrefix(part, "{") {
			return false
		}
	}
	return true
}
