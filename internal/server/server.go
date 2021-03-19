package server

import (
	"safetynet/internal/constants"
)

func Run() {
	r := httpInit()
	r.Run(constants.PORT)
}
