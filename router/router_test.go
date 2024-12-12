package router_test

import (
	request "go-http-server/request"
	"go-http-server/response"
	"go-http-server/router"
	"testing"
)

func TestRouterMathcing(t *testing.T) {
	r := router.NewRouter()
	r.AddHandlerFunc("/test", func(req request.Request) response.Response {
		return response.NewOkTextPlainResponse("test passed")
	})

	req := request.Request{Path: "/test"}
	res := r.Handle(req)

	if res.Body != "test passed" {
		t.Errorf("Expected response body 'test passed', go '%s'", res.Body)
	}

}

func TestDynamicRouteMatching(t *testing.T) {
	r := router.NewRouter()
	r.AddHandlerFunc("/user/{id}", func(req request.Request) response.Response {
		return response.NewOkTextPlainResponse("dynamic route passed")
	})

	req := request.Request{Path: "/user/123"}
	res := r.Handle(req)

	if res.Body != "dynamic route passed" {
		t.Errorf("Expected response body 'dynamic route passed', got '%s'", res.Body)
	}
}
