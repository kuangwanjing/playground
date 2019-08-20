package main

import (
//"time"
)

func main() {
	stopChan := make(chan struct{})
	stoppedChan := make(chan struct{})

	go func() {
		defer func() {
			close(stoppedChan)
			println("stopped")
		}()
		println("start")

		for {
			select {
			case <-stopChan:
				println("hello")
				// stop
				return
			}
		}
	}()

	close(stopChan)
	<-stoppedChan

	//time.Sleep(5 * time.Second)
}
