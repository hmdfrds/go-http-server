package server

import (
	request "go-http-server/Request"
	response "go-http-server/Response"
	"go-http-server/router"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	Router *router.Router
}

func (s *Server) Serve(conn net.Conn) {
	defer conn.Close()

	// Read the incomming
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		log.Printf("Error reading: %v", err)
		return
	}

	// Parse the request
	req, err := request.GetRequest(string(buffer[:n]))
	if err != nil {
		res := response.NewBadRequestResponse()
		conn.Write([]byte(res.String()))
		logRequest(req, res, 0)
		return
	}

	// Check for the requested method and path
	start := time.Now()
	res := s.Router.Handle(req)
	conn.Write([]byte(res.String()))

	logRequest(req, res, time.Since(start))
}

func logRequest(req request.Request, res response.Response, duration time.Duration) {
	log.Printf("Method: %s, Path: %s, Status: %d(%s), Duration: %d ms", req.Method, req.Path, res.StatusCode, http.StatusText(res.StatusCode), duration.Milliseconds())
}
