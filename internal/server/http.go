package server

import (
	"net/http"
	"safetynet/internal/server/handlers"

	"github.com/ChristianStefaniw/cgr"
)

func http_init() *cgr.Router {
	router := cgr.NewRouter()

	router.Route("/alert").Handler(handlers.AlertHandler).Method("POST").Insert()

	router.Route("/viewroutes").Handler(func(w http.ResponseWriter, r *http.Request) {
		for _, route := range router.ViewRouteTree() {
			w.Write([]byte(route))
		}
	}).Method("GET").Insert()

	return router
}
