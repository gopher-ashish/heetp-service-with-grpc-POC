package http

import (
	"github.com/client/service"

	"github.com/gorilla/mux"
)

type ClientHanlder struct {
	*mux.Router
	clientService service.ClientServiceInterface
}

func (h ClientHanlder) Setup(r *mux.Router) {
	r.HandleFunc("/v1/user/create/", h.addUser).Methods("POST")
	r.HandleFunc("/v1/user/", h.getUser).Methods("GET")
}

func NewClientHandler(clinetService service.ClientServiceInterface) ClientHanlder {
	r := mux.NewRouter()

	handler := ClientHanlder{r, clinetService}

	handler.Setup(r)

	return handler
}
