package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	if false {
		ch <- 1
	}
	go func(ch chan int) {
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
