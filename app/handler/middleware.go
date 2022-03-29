package handler

import "net/http"

// JSONContentTypeMiddleware will add the json content type header for all endpoints
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json; charset=UTF-8")
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Add("Access-Control-Allow-Credentials", "true")

		next.ServeHTTP(w, r)
	})
}
