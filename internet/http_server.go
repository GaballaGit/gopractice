package main

import(
	"net/http"
	"fmt"
	"log"
)

func main() {

	// Runs everytime this route is accessed 
	http.HandleFunc("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(resp, "Hello server")
	})

	const serverAddr string = "127.0.0.1:3000"

	// You can also just provide :PORT and go assumes you will use local host
	fmt.Println("Running on port 3000")
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalln("error starting server:", err)
		return
	}

}
