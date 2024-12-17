package calculator

import "fmt"

/*
Public => names starting with Uppercase
Private => name starting with lowercase
*/

// private
/*
var opCount int

func OpCount() int {
	return opCount
}
*/

var opCount map[string]int

func init() {
	fmt.Println("[calc.go] calculator package initialized")
	opCount = make(map[string]int)
}

func OpCount() map[string]int {
	return opCount
}
