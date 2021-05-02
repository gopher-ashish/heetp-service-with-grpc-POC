package service

import (
	"context"

	"github.com/model"
)

type ClientServiceInterface interface {
	CreateUser(context.Context, UserForm) (NewUserReponse, error)
	GetUsers(context.Context) ([]model.User, error)
}
