package main

import (
	"fmt"

	// "github.com/tkmagesh/cisco-gofoundation-dec-2024/11-modularity-demo/calculator"

	// package alias
	calc "github.com/tkmagesh/cisco-gofoundation-dec-2024/11-modularity-demo/calculator"

	// nested package
	"github.com/tkmagesh/cisco-gofoundation-dec-2024/11-modularity-demo/calculator/utils"

	// using OSS package
	"github.com/fatih/color"
)

func main() {
	run()
	color.Red("Hello, World!")
	/*
		calculator.Add(100, 200)
		calculator.Subtract(100, 200)
		fmt.Println("Operation Count :", calculator.OpCount())
	*/

	// using package alias
	calc.Add(100, 200)
	calc.Subtract(100, 200)
	fmt.Println("Operation Count :", calc.OpCount())
	fmt.Println("Is 21 Odd :", utils.IsOdd(21))
}
