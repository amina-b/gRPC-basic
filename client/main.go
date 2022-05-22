package main

import (
	"context"
	"io"
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

	// ValidateEmail function is example of client streaming
	err = ValidateEmail(conn)

	if err != nil {
		log.Fatal(err)
	}

	// GetCourses function is example of server streaming
	err = GetCourses(conn)

	if err != nil {
		log.Fatal(err)
	}

}

func GetCourses(conn *g.ClientConn) error {
	user := PopulateUsers()[0]

	usersClient := m.NewUsersServiceClient(conn)

	resp, err := usersClient.GetCourses(context.Background(), user)

	if err != nil {
		log.Printf("Failed to get courses. Error :%v", err)
		return err
	}

	courses := make([]string, 0)

	for {
		course, err := resp.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Failed to receive courses. Error :%v", err)
		}

		courses = append(courses, course.Course)
	}

	log.Println(courses)

	return nil

}

func ValidateEmail(conn *g.ClientConn) error {

	usersClient := m.NewUsersServiceClient(conn)

	stream, err := usersClient.ValidateUsers(context.Background())

	if err != nil {
		log.Println("Error while creating stream: ", err)
		return err
	}

	requests := PopulateUsers()

	for _, req := range requests {
		err = stream.Send(req)

		if err != nil {
			log.Println("Error while sending stream", err)
			return err
		}
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Println("Error while receiving stream", err)
		return err
	}

	log.Println("Invalid emails are:", resp.InvalidEmail)

	return nil
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
