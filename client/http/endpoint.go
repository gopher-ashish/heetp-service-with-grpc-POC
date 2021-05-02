package http

import (
	"net/http"

	"github.com/client/service"
	"github.com/internal/util"
)

func (h ClientHanlder) addUser(w http.ResponseWriter, req *http.Request) {
	user, err := service.GetUserForm(req)
	if err != nil {
		util.Renderer(w).Error(http.StatusBadRequest, err.Error())
		return
	}

	u, err := h.clientService.CreateUser(req.Context(), *user)
	if err != nil {
		util.Renderer(w).Error(http.StatusBadRequest, err.Error())
		return
	}

	util.Renderer(w).JSON(u)
	return
}

func (h ClientHanlder) getUser(w http.ResponseWriter, req *http.Request) {
	u, err := h.clientService.GetUsers(req.Context())
	if err != nil {
		util.Renderer(w).Error(http.StatusBadRequest, err.Error())
		return
	}

	util.Renderer(w).JSON(u)
	return
}
