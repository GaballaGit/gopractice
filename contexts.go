package main

import(
	"context"
	"fmt"
)

func main() {
	todoContext := context.TODO()
	backgroundContext := context.Background()

	ctx := context.WithValue(todoContext, "name", "Jill")
	fmt.Println(ctx.Value("name"))

	ctx1 := context.WithValue(backgroundContext, "city", "Chicago")
	fmt.Println(ctx1.Value("city"))

}
