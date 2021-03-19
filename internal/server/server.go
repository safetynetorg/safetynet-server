package server

import (
	"safetynet/internal/constants"
)

func Run() {
	r := http_init()
	r.Run(constants.PORT)
}
