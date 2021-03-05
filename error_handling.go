package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by 0! Try other value except 0")
	} else {
		return a / b, nil
	}

}

func main() {
	fmt.Println("Enter First Number: ")

	// var then variable name then variable type
	var a int

	// Taking input from user
	fmt.Scanln(&a)
	fmt.Println("Enter Second Number: ")
	var b int
	fmt.Scanln(&b)
	if r, err := divide(a, b); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(r)
	}
}
