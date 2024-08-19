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

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	request, err := parseRequest(conn)
	if err != nil {
		log.Println("Error Parsing Request,", err)
		return
	}
	
	log.Printf("Request recieved: %s\t%s\t%s", request.Method, request.Path, request.Proto)
	
	response := "HTTP/1.1 200 OK\r\n" +	
				"Content-Type: text/plain\r\n" +
				"Content-Length: 12\r\n" +
				"\r\n" +
				"Hello, World!"
				
	conn.Write([]byte(response))
}