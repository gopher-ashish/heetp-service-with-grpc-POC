package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/PersonalGithubAccount/http-service-with-grpc-POC/protopb"
	"google.golang.org/grpc"
)

type server struct {
	protopb.UnimplementedSumServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to liesten: %v", err)
	}

	s := grpc.NewServer()

	protopb.RegisterSumServer(s, &server{})

	fmt.Println("Started GRPC server....!")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server %v", err)
	}

}

// WelcomeEmail returns welcome
func (*server) WelcomeEmail(ctx context.Context, req *protopb.EmailRequest) (*protopb.EmailResponse, error) {
	email := req.GetEmail()
	log.Print("received request for ", email)
	newMessage := fmt.Sprintf("your new account has been created using email %s. Welcome!", email)
	return &protopb.EmailResponse{Email: newMessage}, nil
}
