package main

import (
	"errors"
	"fmt"
)

func main() {

	for i := 0; i < 20; i++ {
		if _, err := test(i); err != nil {
			fmt.Println(err)
		}
		if _, err := test2(i); err != nil {
			if ae, ok := err.(*errorType); ok {
				fmt.Println(ae.errorMessage)
				fmt.Println(ae.input)
			}
		}
	}

}
func test2(arg int) (int, error) {
	if arg == 12 {
		return -1, &errorType{arg, "this is shit"}

	}
	return arg * 2, nil
}
func test(arg int) (int, error) {
	if arg == 11 {
		return -1, errors.New("this is not acceptable")
	}
	return arg * 2, nil
}

type errorType struct {
	input        int
	errorMessage string
}

func (e *errorType) Error() string {
	return fmt.Sprintf("%d - %s", e.input, e.errorMessage)
}
