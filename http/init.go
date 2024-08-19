package http1

import (
	"fmt"
	"log"
	"net"
)

func Init(port string) error{
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("Error starting server:", err)
	}
	defer listener.Close()
	log.Printf("Starting Server at Port: %s", port)
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection :", err)
			continue
		}
		
		go HandleConnection(conn)
	}
}