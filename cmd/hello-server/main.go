package main

import (
	"fmt"
	"net/http"

	"hello/api"

	_ "hello/docs"
)

func main() {
	http.HandleFunc("/hello", api.HelloServer)
	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server stopped")
}
