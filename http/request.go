package http1

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type Request struct {
	Method string
	Path string
	Proto string
}

func parseRequest(conn net.Conn) (*Request, error) {
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	
	parts := strings.Split(line, " ")
	if len(parts) < 3 {
		return nil, fmt.Errorf("ivalid request line")
	}
	
	return &Request{
		Method: parts[0],
		Path: parts[1],
		Proto: parts[2],
	}, nil
}

// HandleConnection processes the incoming connection and sends a response
func  (s *Server) HandleConnection(conn net.Conn) {
	defer conn.Close()
	
	request, err := parseRequest(conn)
	if err != nil {
		log.Println("Error Parsing Request,", err)
		return
	}
	
	log.Printf("Request recieved: %s\t%s\t%s", request.Method, request.Path, request.Proto)
	
	// Get the handler for the request path
	handler, found := s.handlers[request.Path]
	if !found {
		handler = func(request *Request) *Response {
			return NewTextResponse(4040, "NOT FOUND")
		}
	}
	
	// Generate the response using the handler
	responose := handler(request)
	if responose == nil {
		responose = NewTextResponse(500, "INTERNAL SERVER ERROR")
	}
	
	_, err = conn.Write([]byte(responose.ToString()))
	if err != nil {
		log.Println("Error sending Response:", err)
	}
}