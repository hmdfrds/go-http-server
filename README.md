# Go HTTP Server  

The most *basic* HTTP server written in Go

## Description

This project is for me to explore HTTP servers and improve my Go.

## Features

- Lightweight HTTP server
- Basic routing
- Simple response handling
- Middleware

## Getting Started

### Prequisites

- Go 1.23+

### Installation

1. Clone The Project

    ```bash
    git clone https://github.com/hmdfrds/go-http-server.git
    cd go-http-server
    ```

2. Run The Server

    ```bash
    go run .
    ```

3. Access The Server

    ```bash
    http://localhost:8080
    ```

## To-Do

- [x] Serve static HTML files.
- [ ] Respond with custom messages for different routes.
- [ ] Handle query parameters.
- [x] Return a `404` error for unknown routes.
- [x] Basic logging.
- [ ] Handle `POST` method.

## License

MIT License. See [License](LICENSE).
