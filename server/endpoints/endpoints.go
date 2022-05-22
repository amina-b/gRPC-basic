package endpoints

import (
	"io"
	"log"

	m "github.com/amina-b/gRPC-basic/models"
)

type Server struct {
	m.UnimplementedUsersServiceServer
}

func (s Server) ValidateUsers(stream m.UsersService_ValidateUsersServer) error {

	var invalidEmails []string

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		// Check email validation
		if !isEmailValid(req.Email) {
			invalidEmails = append(invalidEmails, req.Email)
		}
	}

	err := stream.SendAndClose(&m.UserResponse{
		InvalidEmail: invalidEmails,
	})

	if err != nil {
		log.Println("Failed to send and close. Error: ", err)
		return err
	}

	return nil

}

func (s Server) GetCourses(req *m.UserRequest, us m.UsersService_GetCoursesServer) error {

	courses := []string{"Math", "English", "History", "Chemistry"}

	for _, course := range courses {

		err := us.Send(&m.UserCourse{Course: course})

		if err != nil {
			log.Printf("Failed to send course. Error: %v", err)
			return err
		}
	}

	return nil

}
