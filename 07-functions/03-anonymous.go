package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi!")
	}()

	func(userName string) {
		fmt.Printf("Hello %q!\n", userName)
	}("Magesh")

	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	msg := func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
	}("Suresh", "Kannan")
	fmt.Println(msg)

	q, r := func(x, y int) (quotient, remainder int) { //memory allocated & initialized for quotient & remainder
		quotient = x / y
		remainder = x % y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
