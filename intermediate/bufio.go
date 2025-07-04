package intermediate 

import(
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter a sentence")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = input[:len(input)-1]
	fmt.Printf("%s What a brilliant sentence!\n", input)



}
