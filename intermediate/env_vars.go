package intermediate 

import(
	"fmt"
	"os"
	_ "strings"
)

func main() {
	name := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User:", name)
	fmt.Println("Home:", home)

	/*for _, elm := range os.Environ() {
		keypair := strings.SplitN(elm, "=", 2)
		fmt.Println("Key:", keypair[0], "=", keypair[1])
	}*/

	test := os.Getenv("TESTAPI")
	fmt.Println(test)

	os.Setenv("GOPHER", "GO GO")
	fmt.Println(os.Getenv("GOPHER"))
}
