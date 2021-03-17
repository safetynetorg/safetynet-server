package main

import (
	"github.com/ChristianStefaniw/cgr"
	"net/http"
)

const (
	port = "8080"
)

func main() {
	r := cgr.NewRouter()
	r.Route("/").Handler(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("home")) }).Method("GET").Insert()
	r.Run(port)
}
