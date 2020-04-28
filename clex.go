package main

import (
	"hello/client"
	"hello/client/hello_tag"
	"log"
)

func main() {
	cl := client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost("localhost:8080"))
	ep, err := cl.HelloTag.IDOfHelloEndpoint(hello_tag.NewIDOfHelloEndpointParams())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ep)
}
