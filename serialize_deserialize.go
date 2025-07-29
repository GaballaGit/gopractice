package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	alice := Person{
		Name: "Alice",
		Age:  24,
	}

	fmt.Println(alice)
	aliceJson, err := json.Marshal(alice)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(aliceJson))

	var user1 Person
	err = json.Unmarshal(aliceJson, &user1)
	if err != nil {
		panic(err)
	}

	fmt.Println(user1)


	bobby := `{"name": "Bobby", "age": 22}`
	bobbyReader := strings.NewReader(bobby)
	jsonBobby := json.NewDecoder(bobbyReader)

	var user2 Person
	err = jsonBobby.Decode(&user2)
	if err != nil {
		panic(err)
	}

	fmt.Println(user2)
}
