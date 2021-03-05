package main

import (
	"errors"
	"fmt"
	"regexp"
)

/* The password validation criteria is
   Length greater than 7
   At least one upper case character
   At least one special character
   At least one number
*/

func validate(a string) (string, error) {
	if len(a) < 8 {
		return a, errors.New("Password length should be greater than 7")
	}

	num := `[0-9]{1}`
	upper := `[A-Z]{1}`
	special := `[!@#~$%^&*()+|_]{1}`

	if b, err := regexp.MatchString(num, a); !b || err != nil {
		return a, errors.New("Password needs number")
	}

	if b, err := regexp.MatchString(upper, a); !b || err != nil {
		return a, errors.New("Password needs upper case characters")
	}
	if b, err := regexp.MatchString(special, a); !b || err != nil {
		return a, errors.New("Password needs special case characters")
	}
	return a, nil
}

func main() {
	var a string
	fmt.Println("Enter password: ")
	fmt.Scanln(&a)
	if r, err := validate(a); err != nil {
		fmt.Println("Validation error: ", err)
	} else {
		fmt.Println("Correct password: ", r)
	}
}
