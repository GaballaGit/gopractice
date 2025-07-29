package main

import(
	"fmt"
	"net/http"
)

func main() {

	const PORT = 3000

	http.HandleFunc("/hosts", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I- am local host.")
	})


	fmt.Println("server now listening on port:", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
