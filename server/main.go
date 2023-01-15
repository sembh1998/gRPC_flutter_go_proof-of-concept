package main

import (
	"log"

	"github.com/sembh1998/gRCP_flutter_go_proof-of-concept/tree/main/server/src/bootstrap/database"
	"github.com/sembh1998/gRCP_flutter_go_proof-of-concept/tree/main/server/src/bootstrap/server"
)

func main() {
	log.Printf("Hello, world.")

	server.Funca()
	database.Funca()
}
