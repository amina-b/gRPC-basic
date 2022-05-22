package grpc

import (
	"log"
	"net"

	m "github.com/amina-b/gRPC-basic/models"

	e "github.com/amina-b/gRPC-basic/server/endpoints"
	g "google.golang.org/grpc"
)

func GrpcService(port, address string) error {

	listener, err := net.Listen("tcp", address+port)

	if err != nil {
		log.Printf("Failed to create connection. Error: ", err)
		return err
	}

	defer listener.Close()

	s := e.Server{}

	grpcServer := g.NewServer()

	log.Printf("listening on port %v", port)

	m.RegisterGreeterServer(grpcServer, &s)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Printf("Failed to serve. Error :%v", err)
		return err
	}

	return nil

}
