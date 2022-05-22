package endpoints

import (
	"io"
	"log"
	"regexp"

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

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
