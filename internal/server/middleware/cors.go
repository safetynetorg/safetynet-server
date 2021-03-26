package middleware

import (
	"net/http"
)

func CorsMiddleware(w http.ResponseWriter, r *http.Request) {

	// change the url to "localhost:3000" if you're running locally
	w.Header().Add("Access-Control-Allow-Origin", "https://safetynetorg.site")

	w.Header().Add("Access-Control-Allow-Headers", "*")
}
