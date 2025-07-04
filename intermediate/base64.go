package intermediate 

import (
	"encoding/base64"
	"fmt"
)

func main() {

	// Encoding is used for binary transmition as well as data storage



	data := []byte("Hello~, this is a very important message!")

	// Encoding
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decoding
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(decoded)
	fmt.Printf("%s\n", decoded)

	// Url Safe
	urlEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println(urlEncoded)

	urlSafe, err := base64.URLEncoding.DecodeString(urlEncoded)
	if err != nil {
		panic(err)
	}

	fmt.Println(urlSafe)
	fmt.Printf("%s\n", urlSafe)
}
