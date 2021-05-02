package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"

	clientHandler "github.com/PersonalGithubAccount/http-service-with-grpc-POC/client/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type appHandler func(http.ResponseWriter, *http.Request) error

// in a consistent manner
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Remove in production
	devMode := true
	if devMode {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Info(err.Error())
		} else {
			log.Info(string(dump))
		}
	}

	if err := fn(w, r); err != nil { // err is *model.Err, not os.Err.
		json.NewEncoder(w).Encode(err)
	}
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func AttachRoutes(client clientHandler.ClientHanlder) http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	client.Setup(r)

	r.PathPrefix("/").HandlerFunc(catchAllHandler).Methods("GET")

	return r
}
