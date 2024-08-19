# Basic HTTP/1.1 Server in Go

This Go package implements a basic HTTP/1.1 server. It demonstrates how to set up a server, handle connections, and serve HTTP requests.

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)

## Overview

This package provides a simple implementation of an HTTP/1.1 server in Go. The server listens on a specified port, accepts incoming connections, and handles HTTP requests according to the HTTP/1.1 specification.

## Installation

To use this package, you need Go installed on your system. You can install it from the [this repo]([https://golang.org/dl/](https://github.com/high-horse/http-1.1-implementation-in-go/).

Clone this repository and build the project:

```bash
git clone https://github.com/high-horse/http-1.1-implementation-in-go
cd http-1.1-implementation-in-go
make build
```

## Usage

To start the server, you can use the server executable or import the package into your Go project.
Running the Server

Run the server with the following command:

```bash
make serve
```

## Code Example

Here’s a basic example of how to use the package in your own Go application:

```bash
package main

import(
	"log"
	
	ht"http1.1/http"
)

func main() {
  // Instanciate new server
	server := ht.NewServer()

  // Register handler functions
	server.Handle("/json", handleJSON)
	server.Handle("/string", handleString)

  //  start the server
	err := server.Init(":8000") 
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// handleString handles requests to the root path ("/string")
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

```
