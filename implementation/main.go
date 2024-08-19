package main

import(
	"log"
	
	ht"http1.1/http"
)

func main() {
	server := ht.NewServer()
	
	server.Handle("/", func(request *ht.Request) *ht.Response {
		return ht.NewTextResponse(200, "Hello, World!")
	})
	
	server.Handle("/json", func(request *ht.Request) *ht.Response {
		data := map[string]string{"message": "Hello, JSON!"}
		response, err := ht.NewJSONResponse(200, data)
		if err != nil {
			return ht.NewTextResponse(500, "Internal Server Error")
		}
		return response
	})
	server.Handle("/jsonF", handleJSON)
	server.Handle("/string", handleString)
	
	err := server.Init(":8000") 
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// handleString handles requests to the root path ("/")
func handleString(r *ht.Request) *ht.Response {
	return ht.NewTextResponse(200, "Hello, World!")
}

// handleJSON handles requests to the "/json" path
func handleJSON(request *ht.Request) *ht.Response {
	data := map[string]string{"message": "Hello, JSON!"}
	response, err := ht.NewJSONResponse(200, data)
	if err != nil {
		return ht.NewTextResponse(500, "Internal Server Error")
	}
	return response
}
