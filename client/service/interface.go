package service

import (
	"context"

	"github.com/PersonalGithubAccount/http-service-with-grpc-POC/model"
)

type ClientServiceInterface interface {
	CreateUser(context.Context, UserForm) (NewUserReponse, error)
	GetUsers(context.Context) ([]model.User, error)
}
