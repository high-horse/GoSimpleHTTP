.PHONY: server build serve curl-call help

# Default target
all: help

# Run the server directly from the source code
server:
	cd cmd && go run .

# Build the project
build:
	go build -o server cmd/main.go

# Serve from the built binary
serve: build
	./server

# Make a curl call to the server
curl-call:
	curl -v http://localhost:8000/string

# Display help messages
help:
	@echo "Usage:"
	@echo "  make server       - Run the server directly from source code"
	@echo "  make build        - Build the server binary"
	@echo "  make serve        - Build the server binary and run it"
	@echo "  make curl-call    - Make a curl call to the running server"
	@echo "  make help         - Show this help message"

