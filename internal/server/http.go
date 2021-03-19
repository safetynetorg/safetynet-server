package server

import (
	"net/http"
	"safetynet/internal/server/handlers"

	"github.com/ChristianStefaniw/cgr"
)

func httpInit() *cgr.Router {
	router := cgr.NewRouter()

	router.Route("/alert").Handler(handlers.FindDevicesToAlert).Method("POST").Insert()
	router.Route("/updatelocation").Handler(handlers.UpdateLocation).Method("PUT").Insert()
	router.Route("/check").Handler(handlers.Alert).Method("GET").Insert()

	router.Route("/viewroutes").Handler(func(w http.ResponseWriter, r *http.Request) {
		for _, route := range router.ViewRouteTree() {
			w.Write([]byte(route))
		}
	}).Method("GET").Insert()

	return router
}
