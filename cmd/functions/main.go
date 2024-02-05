package main

import (
	"errors"
	"fmt"
)

func main() {

	printMessage("Hello World")

	var result, remainder, err = divideNumbers(101, 10)

	if err != nil {
		fmt.Printf(err.Error())
	} else {

		fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)
	}

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("Result: %d\n", result)
	default:
		fmt.Println("Result: ", result, "Remainder: ", remainder)
	}

	// conditional switch statement
	switch remainder {
	case 0:
		fmt.Printf("Result: %d\n", result)
	case 1, 2:
		fmt.Println("Result: ", result, "Remainder: ", remainder)
	default:
		fmt.Println("Result: ", result, "Remainder: ", remainder)
	}

}

func printMessage(msg string) {
	fmt.Println(msg)
}

func divideNumbers(num1 int, num2 int) (int, int, error) {
	// by default error value is nil
	var err error
	if num2 == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result = num1 / num2
	var remainder = num1 % num2
	return result, remainder, err
}
