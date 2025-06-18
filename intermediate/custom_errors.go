package intermediate

import (
	"errors"
	"fmt"
)

type customError struct {
	code    int
	message string
	er      error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s %v", e.code, e.message, e.er)
}

func something() error {
	return &customError{
		code:    500,
		message: "Bad gateway request!",
	}
}

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("Internal error")
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}

}
