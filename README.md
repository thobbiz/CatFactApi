## CatFactApi

CatFactApi is a backend service built with Golang that sends requests to get random catfacts.

## Features:
- rate limiting using `golang.org/x/time/rate`
- error handling

## Stack and Tools:
- Golang
- Gin Web Framework
- net/http
- Go Rate Limiting

## Get Started
### Prerequisites
- G0 1.21+
- Internet Connection
- Modules:
  ```bash
  go get github.com/gin-gonic/gin
  go get golang.org/x/time/rate

## Usage
- Start the server:
  ```bash
  go run main.go
- Send a request:
  ```bash
  curl http://localhost:7070/me
