package main

import(
	ht"http1.1/http"
)

func main() {
	err := ht.Init(":8000")
	if err != nil {
		println("error starting server", err)
	}
}