package main

import (
	"log"

	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/bootstrap/server/grpcapi"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	server, err := grpcapi.NewServer()
	if err != nil {
		log.Fatal("cannot create grpc server", err)
	}
	server.Initialice(config)
}
