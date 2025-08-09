package main

import(
	"net/http"
	"fmt"
	"io"
)

func main() {
	client := &http.Client{}

	resp, err := client.Get("http://127.0.0.1:3000/")
	if err != nil {
		fmt.Println("error making get request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print response Body

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return
	}

	fmt.Println(string(body))

}
