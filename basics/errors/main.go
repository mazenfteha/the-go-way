package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(doSomething())

	result, err := divide(10, 0)
	if err != nil { // Error Checking
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)

	var NotFoundError = errors.New("not found")
	err = fmt.Errorf("an error occurred: %w", NotFoundError)
	if errors.Is(err, NotFoundError) {
		fmt.Println("The error is 'not found'")
	}
}

func doSomething() error {
	return errors.New("Something went wrong")
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero", a)
	}
	return a / b, nil
}

/*
Custom Error Types
*/
type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Wrapping and Unwrapping Errors
