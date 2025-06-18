package intermediate

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

	John.changeName("Bob")

	fmt.Println("Actually, my name is now " + John.firstName + ".\nOh! I see.")

}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) changeName(newName string) {
	p.firstName = newName
}
