package server

import (
	"github.com/ChristianStefaniw/cgr"
	"net/http"
	"safetynet/internal/handlers"
	"safetynet/internal/middleware"
)

func httpInit() *cgr.Router {
	router := cgr.NewRouter()
	corsMiddleware := cgr.NewMiddleware(middleware.CorsMiddleware)
	corsPreflightConf := cgr.NewRouteConf()
	corsPreflightConf.HandlePreflight(true)

	// api endpoints
	router.Route("/").Handler(handlers.Home).Method("GET").Insert()
	router.Route("/alert").Handler(handlers.FindDevicesToAlert).Method("POST").Insert()
	router.Route("/new").Handler(handlers.NewDevice).Method("POST").Insert()
	router.Route("/signup").Handler(handlers.SignUp).Assign(corsMiddleware).SetConf(corsPreflightConf).Method("POST", "OPTIONS").Insert()
	router.Route("/contact").Handler(handlers.Contact).Assign(corsMiddleware).SetConf(corsPreflightConf).Method("POST", "OPTIONS").Insert()
	router.Route("/updatelocation").Handler(handlers.UpdateLocation).Method("PUT").Insert()
	router.Route("/delete/:id").Handler(handlers.DeleteDevice).Method("DELETE").Insert()

	router.Route("/viewroutes").Handler(func(w http.ResponseWriter, r *http.Request) {
		for _, route := range router.ViewRouteTree() {
			w.Write([]byte(route))
		}
	}).Method("GET").Insert()

	return router
}
