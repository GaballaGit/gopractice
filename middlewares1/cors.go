package middlewares1

import (
	"fmt"
	"net/http"
)

var allowedOrigins = []string{
	"https://127.0.0.1:8080",
	"https://my-example-url",
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("origin")
		fmt.Println("origin:", origin)
		if !isOriginAllowed(r.Header.Get("origin")) {
			http.Error(w, "Origin blocked by CORS", http.StatusForbidden)
			return
		}

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "3600")

		next.ServeHTTP(w, r)
	})

}

func isOriginAllowed(domain string) bool {
	for _, elm := range allowedOrigins {
		if elm == domain {
			return true
		}
	}
	return false
}
