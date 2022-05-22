package main

import (
	"context"
	"log"

	m "github.com/amina-b/gRPC-basic/models"

	g "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := g.DialContext(context.Background(), "localhost:2000", g.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to create connection. Error: %v", err)
	}

	defer conn.Close()

	greaterClient := m.NewGreeterClient(conn)

	resp, err := greaterClient.SayHello(context.Background(), &m.HelloRequest{})

	if err != nil {
		log.Println("err", err)
	}

	log.Println(resp)
}
