package main

import (
	"errors"
	"fmt"
)

func main() {
	divisor := 0
	for {
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		if q, r, err := divide(100, divisor); err != nil {
			fmt.Println("Error :", err)
			fmt.Println("Try again!")
			continue
		} else {
			fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
			break
		}
	}
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divide by zero error")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("divide by zero error")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
