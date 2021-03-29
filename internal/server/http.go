package server

import (
	"github.com/ChristianStefaniw/cgr"
	"net/http"
	"safetynet/internal/server/handlers"
	"safetynet/internal/server/middleware"
)

func httpInit() *cgr.Router {
	router := cgr.NewRouter()
	corsMiddleware := cgr.NewMiddleware(middleware.CorsMiddleware)

	// api endpoints
	router.Route("/").Handler(handlers.Home).Method("GET").Insert()
	router.Route("/alert").Handler(handlers.FindDevicesToAlert).Method("POST").Insert()
	router.Route("/new").Handler(handlers.NewDevice).Method("POST").Insert()
	router.Route("/signup").Handler(handlers.SignUp).Assign(corsMiddleware).Method("POST").Insert()
	router.Route("/contact").Handler(handlers.Contact).Assign(corsMiddleware).Method("POST").Insert()
	router.Route("/updatelocation").Handler(handlers.UpdateLocation).Method("PUT").Insert()
	router.Route("/delete/:id").Handler(handlers.DeleteDevice).Method("DELETE").Insert()

	// cors preflight
	corsPreflight(router, corsMiddleware)

	router.Route("/viewroutes").Handler(func(w http.ResponseWriter, r *http.Request) {
		for _, route := range router.ViewRouteTree() {
			w.Write([]byte(route))
		}
	}).Method("GET").Insert()

	return router
}

func corsPreflight(router *cgr.Router, middleware *cgr.Middleware) {

	router.Route("/signup").Handler(ok).Assign(middleware).Method("OPTIONS").Insert()
	router.Route("/contact").Handler(ok).Assign(middleware).Method("OPTIONS").Insert()
}

func ok(rw http.ResponseWriter, r *http.Request) { rw.WriteHeader(http.StatusOK) }
