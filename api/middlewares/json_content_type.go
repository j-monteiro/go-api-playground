package middlewares

import (
	"fmt"
	"net/http"
)

func JSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Setting Content-Type to application/json")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
