package basics 

import "fmt" 

func main() {
	process()
	fmt.Println("Recover from panic")
}

func process() {
	defer func() {
	if r := recover(); r != nil {
		fmt.Println("Recoverd: ", r)
		} 
	} ()
	
	panic("Oh no")
}
