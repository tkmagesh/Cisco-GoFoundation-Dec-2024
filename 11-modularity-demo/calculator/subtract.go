package calculator

import "fmt"

func init() {
	fmt.Println("[subtract.go] calculator package initialized")
}

// public
func Subtract(x, y int) {
	opCount["Subtract"]++
	fmt.Println("Subtract Result :", x-y)
}
