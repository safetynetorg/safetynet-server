package server

import (
	"safetynet/internal/controllers"
	"safetynet/internal/middleware"

	"github.com/ChristianStefaniw/cgr-v2"
)

func httpInit() *cgr.Router {
	router := cgr.NewRouter()
	corsMiddleware := cgr.NewMiddleware(middleware.CorsMiddleware)

	// api endpoints
	router.Route("/").Handler(controllers.Home).Method("GET").Insert()
	router.Route("/alert").Handler(controllers.FindDevicesToAlert).Method("POST").Insert()
	router.Route("/new").Handler(controllers.NewDevice).Method("POST").Insert()
	router.Route("/signup").Handler(controllers.SignUp).Assign(corsMiddleware).HandlePreflight().Method("POST", "OPTIONS").Insert()
	router.Route("/contact").Handler(controllers.Contact).Assign(corsMiddleware).HandlePreflight().Method("POST", "OPTIONS").Insert()
	router.Route("/updatelocation").Handler(controllers.UpdateLocation).Method("PUT").Insert()
	router.Route("/delete/:id").Handler(controllers.DeleteDevice).Method("DELETE").Insert()

	return router
}
