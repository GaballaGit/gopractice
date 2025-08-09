package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	mw "gopractice/middlewares1"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	cert := "cert.pem"
	key := "key.pem"

	PORT := os.Getenv("PORT")

	tlsCnf := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	mux := http.NewServeMux()

	secureMux := applyMiddleware(
		mux,
		mw.Cors,
	)
	server := &http.Server{
		Addr:      ":" + PORT,
		Handler:   secureMux,
		TLSConfig: tlsCnf,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Root!\n"))
	})

	fmt.Println("Starting server on port:", PORT)
	err = server.ListenAndServeTLS(cert, key)
}

type middleware func(http.Handler) http.Handler

func applyMiddleware(handler http.Handler, middlewares ...middleware) http.Handler {
	for _, midware := range middlewares {
		handler = midware(handler)
	}
	return handler
}
