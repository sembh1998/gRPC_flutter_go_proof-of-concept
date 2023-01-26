package main

import (
	"log"

	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/bootstrap/server"
)

func main() {
	log.Printf("Hello, world.")

	server, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	server.Initialice()
}
