package main

import (
	"fmt"
	"sync"
)

// communicate by sharing memory
var result int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Add(1)
	go add(300, 400, wg)
	wg.Add(1)
	go add(200, 500, wg)
	wg.Wait()
	fmt.Println("Result :", result)
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	// time.Sleep(2 * time.Second)
	mutex.Lock()
	{
		result += x + y
	}
	mutex.Unlock()
}
