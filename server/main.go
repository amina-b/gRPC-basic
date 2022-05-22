package main

import (
	"log"

	"github.com/amina-b/gRPC-basic/server/config"
	"github.com/amina-b/gRPC-basic/server/grpc"
)

func main() {

	// Load configuration
	err := config.Load()

	if err != nil {
		log.Fatal(err)
	}
	// Start grpc service
	err = grpc.GrpcService(config.Config.Environment.Port, config.Config.Environment.Address)

	if err != nil {
		log.Fatal(err)
	}

}
