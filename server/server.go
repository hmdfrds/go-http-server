package server

import (
	"fmt"
	"go-http-server/handler"
	request "go-http-server/request"
	response "go-http-server/response"
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

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		log.Printf("Error reading: %v", err)
		return
	}

	req, err := request.GetRequest(string(buffer[:n]))
	if err != nil {
		res := response.NewBadRequestResponse()
		conn.Write([]byte(res.String()))
		logRequest(req, res, 0)
		return
	}

	start := time.Now()
	res := s.Router.Handle(req)
	conn.Write([]byte(res.String()))

	logRequest(req, res, time.Since(start))
}

func logRequest(req request.Request, res response.Response, duration time.Duration) {
	log.Printf("Method: %s, Path: %s, Status: %d(%s), Duration: %d ms", req.Method, req.Path, res.StatusCode, http.StatusText(res.StatusCode), duration.Milliseconds())
}

func StartServer() {
	// Create router
	router := router.NewRouter()
	router.AddHandlerFunc("/", handler.ServeHTML)
	router.AddHandlerFunc("/home", handler.ServeHTML)
	router.AddHandlerFunc("/about", handler.ServeHTML)
	router.AddHandlerFunc("/form", handler.ServeHTML)
	router.AddHandlerFunc("/hello/{name}", handler.HelloHandler)
	router.AddHandlerFunc("/static/style.css", handler.StaticFileHandler)
	router.AddHandlerFunc("/search", handler.SearchHandler)
	router.AddHandlerFunc("/submit-form", handler.SubmitFormHandler)

	// Create server
	server := Server{Router: router}

	// Open the TCP socket
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("Listening on port 8080...")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		// Handler connection
		go server.Serve(conn)
	}
}
