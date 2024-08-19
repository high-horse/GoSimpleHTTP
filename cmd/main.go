package main

import(
	"net"
	"fmt"
	
	ht"http1.1/http"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection :", err)
			continue
		}
		
		go ht.HandleConnection(conn)
	}
}