package server

import (
	"safetynet/internal/constants"
)

// run the server
func Run() {
	r := httpInit()
	r.Run(constants.PORT)
}
