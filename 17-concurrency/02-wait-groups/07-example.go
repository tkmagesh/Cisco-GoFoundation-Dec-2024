/*
Modify the below program in such a way that the logic for checing a given number is a prime number is executed concurrently
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var start, end int
	wg := &sync.WaitGroup{}
	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			printIfPrime(no)
		}()
	}
	wg.Wait()
}

func printIfPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
