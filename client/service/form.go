package service

import (
	"net/http"

	"github.com/PersonalGithubAccount/http-service-with-grpc-POC/internal/util"
	"github.com/PersonalGithubAccount/http-service-with-grpc-POC/model"
)

type UserForm struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func GetUserForm(r *http.Request) (*UserForm, error) {
	f := new(UserForm)
	req := util.Request{r}
	if err := req.Decode(&f); err != nil {
		return nil, err
	}

	return f, nil
}

type NewUserReponse struct {
	WelcomeMessage string     `json:"welcomeMessage"`
	User           model.User `json:"user"`
}
