package main

import "fmt"

// MyString is an alias for "string"
type MyString string

func (s MyString) Length() int {
	return len(s)
}

func main() {
	s := MyString("Labore qui ex consequat do ut pariatur fugiat nulla magna veniam aute.")
	fmt.Println(s.Length())
}
