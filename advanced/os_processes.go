package advanced

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// basically does the following:
	// echo "food is good\nbar\az\n" | grep foo
	pr, pw := io.Pipe()
	cmd := exec.Command("grep", "foo")
	cmd.Stdin = pr

	go func() {
		defer pw.Close()
		pw.Write([]byte("food is good\nbar\naz\n"))
	}()

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", string(output))

	/*
		cmd := exec.Command("printenv", "SHELL")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(output))
	*/

	/*
		cmd := exec.Command("sleep", "60")
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(2 * time.Second)

		err = cmd.Process.Kill()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Process killed")

		err = cmd.Wait()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Process is complete")
	*/

	/*
		cmd := exec.Command("grep", "foo")

		cmd.Stdin = strings.NewReader("foo\nbar\naz\n")

		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Output:", string(output))
	*/
	/*
		cmd := exec.Command("echo", "Hello World!")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Output:", string(output))
	*/
}
