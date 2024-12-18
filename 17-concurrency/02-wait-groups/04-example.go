package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1) // increment the wg counter by 1
	go f1(wg)
	f2()
	wg.Wait() // block the execution of this function until the internal counter becomes 0 (default = 0)
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Println("f1 started")
	simulatedTime := rand.Intn(5)
	fmt.Printf("[f1] simulating a time consuming operation for %vs duration\n", simulatedTime)
	time.Sleep(time.Duration(simulatedTime) * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
