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

// communicate by sharing memory
var primes []int

func main() {
	var start, end int
	wg := &sync.WaitGroup{}
	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)
	wg.Add(1)
	go GenPrimes(start, end, wg)
	wg.Wait()
	for _, primeNo := range primes {
		fmt.Println("Prime No :", primeNo)
	}
}

func GenPrimes(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg2 := &sync.WaitGroup{}
	var mutex sync.Mutex
	for no := start; no <= end; no++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			if IsPrime(no) {
				mutex.Lock()
				{
					primes = append(primes, no)
				}
				mutex.Unlock()
			}
		}()
	}
	wg2.Wait()

}

func IsPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
