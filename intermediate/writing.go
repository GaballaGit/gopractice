package intermediate 

import(
	"os"
	"fmt"
)

func main() {

	file, err := os.Create("Output.txt")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()

	data := []byte("This is my outputtttttt\n")

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}




}
