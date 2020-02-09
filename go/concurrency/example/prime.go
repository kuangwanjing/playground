package main

import "fmt"

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)

	for i := 0; i < 20; i++ {
		num := <-ch
		fmt.Println(num)
		nch := make(chan int)
		go filter(ch, nch, num)
		ch = nch
	}
}
