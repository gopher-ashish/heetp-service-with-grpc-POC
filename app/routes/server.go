package routes

import (
	"net/http"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func ServeHTTP(h http.Handler, port string) {
	log.Info("HTTP services loaded and serving on port ", port)

	headers := handlers.AllowedHeaders([]string{"Accept", "Accept-Encoding", "Authorization",
		"Content-Length", "Content-Type", "Origin", "X-CSRF-Token"})

	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"DELETE", "GET", "PATCH", "POST", "PUT"})

	n := negroni.Classic()
	n.UseHandler(handlers.CORS(headers, origins, methods)(
		handlers.CompressHandler(h),
	))

	log.Fatal(http.ListenAndServe(":"+port, n))
}
