package http1

import (
	"fmt"
	"log"
	"net"
)

func NewServer() *Server {
	return &Server{
		handlers: make(map[string]HandlerFunc),
	}
}

func (s *Server)Handle(path string, handler HandlerFunc){
	s.handlers[path] =handler
}

const defaultPort = "8080"

func (s *Server) Init(port string) error{
	if port == "" {
		port = defaultPort
	}
	
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("Error starting server: %w", err)
	}
	defer listener.Close()
	log.Printf("Starting Server at Port: %s", port)
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection :", err)
			continue
		}
		
		go s.HandleConnection(conn)
	}
}

