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

	usersClient := m.NewUsersServiceClient(conn)

	stream, err := usersClient.ValidateUsers(context.Background())

	if err != nil {
		log.Println("Error while creating stream: ", err)
	}

	requests := PopulateUsers()

	for _, req := range requests {
		err = stream.Send(req)

		if err != nil {
			log.Println("Error while sending stream", err)
		}
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Println("Error while receiving stream", err)
	}

	log.Println("Invalid emails are:", resp.InvalidEmail)

}

func PopulateUsers() []*m.UserRequest {
	requests := []*m.UserRequest{
		{Id: 1, Name: "example 1", Email: "example1@gmail.com"},
		{Id: 2, Name: "example 2", Email: "example2@gmail.com"},
		{Id: 3, Name: "example 3", Email: "example"},
		{Id: 4, Name: "example 4", Email: "example4@gmail.com"},
		{Id: 5, Name: "example 3", Email: "example"},
	}

	return requests

}
