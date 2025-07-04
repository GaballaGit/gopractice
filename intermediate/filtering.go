package intermediate

import (
	"bufio"
	_ "crypto/sha512"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Error opening log:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	filter := "[REDACTED]"

	for scanner.Scan() {
		token := scanner.Text()
		if strings.Contains(token, filter) == true {
			fmt.Printf("\033[1;31m%s\033[0m\n", token)
		} else {
			fmt.Println(token)
		}
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}
}
