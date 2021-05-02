package service

import (
	"context"
	"log"

	grpcb "github.com/PersonalGithubAccount/http-service-with-grpc-POC/protopb"

	"github.com/PersonalGithubAccount/http-service-with-grpc-POC/model"
)

type clientService struct {
	grpcClient grpcb.SumClient
	clientRepo model.UserRepository
}

//Create User function create user in databse and send email to grpc service
func (c *clientService) CreateUser(ctx context.Context, f UserForm) (NewUserReponse, error) {
	user, err := c.clientRepo.Create(c.buildUser(f))
	if err != nil {
		return NewUserReponse{}, err
	}

	res, err := c.grpcClient.WelcomeEmail(ctx, &grpcb.EmailRequest{Email: user.Email})
	if err != nil {
		log.Print("Error: ", err)
	}

	return NewUserReponse{User: user, WelcomeMessage: res.String()}, nil
}

//GetUsers function retrives al users from system
func (c *clientService) GetUsers(ctx context.Context) ([]model.User, error) {
	return c.clientRepo.ShowUsers()
}

//new service
func NewClientService(grpcClient grpcb.SumClient, repo model.UserRepository) ClientServiceInterface {
	return &clientService{grpcClient, repo}
}
