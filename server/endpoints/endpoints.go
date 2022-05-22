package endpoints

import (
	"context"
	"log"

	m "github.com/amina-b/gRPC-basic/models"
)

type Server struct {
	m.UnimplementedGreeterServer
}

func (s Server) SayHello(ctx context.Context, r *m.HelloRequest) (*m.HelloReply, error) {
	log.Println("request comming")

	response := new(m.HelloReply)

	response.Message = "helloo from response"

	return response, nil

}
