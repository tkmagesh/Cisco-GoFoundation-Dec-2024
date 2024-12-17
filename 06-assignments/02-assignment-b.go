/*
Create an interactive calculator

Display the following menu
1. Add
2. Subtract
3. Multiply
4. Divide
5. Exit

Accept the user choice
if user choice == 5
	then exit

if user choice == 1 to 4
	accept the 2 operands from the user
	perform the corresponding operation
	print the result
	display the menu again

if user choice < 1 OR > 5
	print "invalid choice"
	display the menu again

*/

package main

import "fmt"

func main() {
	var n1, n2, userChoice, result int
LOOP:
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter the choice:")
		fmt.Scanln(&userChoice)

		switch userChoice {
		case 1:
			fmt.Println("Enter the operands")
			fmt.Scanln(&n1, &n2)
			result = n1 + n2
			fmt.Println("Result :", result)
		case 2:
			fmt.Println("Enter the operands")
			fmt.Scanln(&n1, &n2)
			result = n1 - n2
			fmt.Println("Result :", result)
		case 3:
			fmt.Println("Enter the operands")
			fmt.Scanln(&n1, &n2)
			result = n1 * n2
			fmt.Println("Result :", result)
		case 4:
			fmt.Println("Enter the operands")
			fmt.Scanln(&n1, &n2)
			result = n1 / n2
			fmt.Println("Result :", result)
		case 5:
			break LOOP
		default:
			fmt.Println("Invalid choice")
			continue
		}
	}
}
