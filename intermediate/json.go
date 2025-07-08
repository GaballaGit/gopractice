package intermediate 

import(
	"fmt"
	"encoding/json"
)

type Person struct {
	Firstname string `json:"name"`
	Age int `json:"age"`
}

func main() {

	john := Person{Firstname: "John", Age: 21}

	data, err := json.Marshal(john)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))


}
