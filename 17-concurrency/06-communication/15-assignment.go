/*
Refactor this as follows

 1. IsPrime(no int) => responsible for checking if the given number is prime number
 2. GenPrimes(start, end)
    => Gen prime numbers CONCURRENTLY
 3. main()
    => Printing the generated prime numbers
    => Execute the GenPrimes() as a goroutine
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
		go printIfPrime(wg, no)
	}
	wg.Wait()
}

func printIfPrime(wg *sync.WaitGroup, no int) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
