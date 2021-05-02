package service

import "github.com/PersonalGithubAccount/http-service-with-grpc-POC/model"

type additionResponse struct {
	Result uint `json:"result"`
}

func (c clientService) buildUser(u UserForm) model.User {
	return model.User{
		Username:  u.Username,
		FirstName: u.Name,
		Email:     u.Email,
	}
}
