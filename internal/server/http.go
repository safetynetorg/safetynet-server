package server

import (
	"net/http"
	"safetynet/internal/pkg/constants"

	"github.com/ChristianStefaniw/cgr"
)

func http_init() {
	r := cgr.NewRouter()
	r.Route("/").Handler(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	}).Method("GET").Insert()
	r.Run(constants.PORT)
}

func Run() {
	http_init()
}
