package server

import "os"

// run the server
func Run() {
	r := httpInit()
	r.Run(os.Getenv("PORT"))
}
