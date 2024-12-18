package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func genNos(ch chan int) {
	for no := range 10 {
		time.Sleep(500 * time.Millisecond)
		ch <- no * 10
	}
	close(ch)
}
