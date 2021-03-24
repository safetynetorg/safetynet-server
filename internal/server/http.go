package server

import (
	"net/http"
	"safetynet/internal/server/handlers"

	"github.com/ChristianStefaniw/cgr"
)

func httpInit() *cgr.Router {
	router := cgr.NewRouter()

	// api endpoints
	router.Route("/alert").Handler(handlers.FindDevicesToAlert).Method("POST").Insert()
	router.Route("/new").Handler(handlers.NewDevice).Method("POST").Insert()
	router.Route("/signup").Handler(handlers.SignUp).Method("POST").Insert()
	router.Route("/updatelocation").Handler(handlers.UpdateLocation).Method("PUT").Insert()
	router.Route("/delete/:id").Handler(handlers.DeleteDevice).Method("DELETE").Insert()

	router.Route("/viewroutes").Handler(func(w http.ResponseWriter, r *http.Request) {
		for _, route := range router.ViewRouteTree() {
			w.Write([]byte(route))
		}
	}).Method("GET").Insert()

	return router
}
