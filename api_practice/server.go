package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	_"golang.org/x/net/http2"
)

func loadClientAs() *x509.CertPool {
	clientAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalf("Could not load clients CA:", err)
	}
	clientAs.AppendCertsFromPEM(caCert)
	return clientAs
}

func main() {


	port := 3000

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handle incoming orders")
		logRequestDetail(r)
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handle incoming users")
		logRequestDetail(r)
	})

	// TLS Cert and Key
	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		//ClientAuth: tls.RequireAndVerifyClientCert,
		//ClientCAs: loadClientAs(),
	}


	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		TLSConfig: tlsConfig,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}


	//http2.ConfigureServer(server, &http2.Server{})



	fmt.Println("Server is listening on port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalf("could not start server: %s", err)
	}

}


func logRequestDetail(r *http.Request) {
	httpVersion := r.Proto
	log.Println("Recieved request with HTTP Version:", httpVersion)
	if r.TLS != nil {
		tlsVer := getTLSVersion(r.TLS.Version)
		log.Println("Made with", tlsVer)
	}
}

func getTLSVersion(version uint16) string {
	switch version{
	case tls.VersionTLS10:
	return "TLS 1.0"
	case tls.VersionTLS11:
	return "TLS 1.1"
	case tls.VersionTLS12:
	return "TLS 1.2"
	case tls.VersionTLS13:
	return "TLS 1.3"
	default:
	return "Unknown TLS Version"
	}
}
