/*
refactor the "interactive calculator" assignment solution using functions
Ensure that the functions follow SRP (Single Responsibility Principle)
*/

package main

import "fmt"

func main() {
	var userChoice int
	for {
		userChoice = getUserChoice()
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		processUserChoice(userChoice)
	}
}

func processUserChoice(userChoice int) {
	var result int
	n1, n2 := getOperands()
	switch userChoice {
	case 1:
		add(n1, n2, &result)
	case 2:
		subtract(n1, n2, &result)
	case 3:
		multiply(n1, n2, &result)
	case 4:
		divide(n1, n2, &result)
	}
	fmt.Println("Result :", result)
}

func add(x, y int, result *int) {
	*result = x + y
}

func subtract(x, y int, result *int) {
	*result = x - y
}

func multiply(x, y int, result *int) {
	*result = x * y
}

func divide(x, y int, result *int) {
	*result = x / y
}

func getOperands() (n1, n2 int) {
	fmt.Println("Enter the operands")
	fmt.Scanln(&n1, &n2)
	return
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter the choice:")
	fmt.Scanln(&userChoice)
	return userChoice
}
