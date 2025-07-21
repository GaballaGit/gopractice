package advanced

import (
	"fmt"
	"reflect"
)

func main() {

	x := 2

	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Println("Value:", v)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", v.Kind())
	fmt.Println("Is int:", v.Kind() == reflect.Int)
	fmt.Println("Is string:", v.Kind() == reflect.String)
	fmt.Println("Is Zero:", v.IsZero())

	y := 10
	v1 := reflect.ValueOf(&y).Elem()
	v2 := reflect.ValueOf(&y)

	fmt.Println("Type of v1:", v1)
	fmt.Println("Type of v1:", v2)

}
