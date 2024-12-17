package main

import "fmt"

func main() {
	/*
		func(string)
		func(int, string)
		func(int, int) int
		func(int, int) (int, string)
		func(...int)int
	*/
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi!")
	}
	sayHi()
}
