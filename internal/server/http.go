package server

import (
	"net/http"
	"safteynet/internal/pkg/constants"

	"github.com/ChristianStefaniw/cgr"
)

func http_init() {
	r := cgr.NewRouter()
	r.Route("/").Handler(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	}).Method("GET").Insert()
	r.Run(constants.Port)
}

func Run() {
	http_init()
}
