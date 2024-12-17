package calculator

import "fmt"

func init() {
	fmt.Println("[add.go] calculator package initialized")
}

// public
func Add(x, y int) {
	opCount["Add"]++
	fmt.Println("Add Result :", x+y)
}
