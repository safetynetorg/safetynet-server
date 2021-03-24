package server

import "safetynet/internal/keys"

// run the server
func Run() {
	r := httpInit()
	r.Run(keys.PORT)
}
