package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	John := Person{
		firstName: "John",
		lastName:  "Golang",
		age:       15,
	}

	fmt.Println("Greetings, " + John.fullName() + ".")

	//Annoynmous struct
	user := struct {
		username string
		age      int
	}{
		username: "Miku",
		age:      20,
	}

	fmt.Println("Why hello " + user.username + ".")
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}
