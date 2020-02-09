// This is an example of the way i used fan-out pattern to let the task executors receive a taks signal and return the
// result back to the initiator.
// but i caused the leaking of goroutine by not waiting for the return of every goroutines.
// And it is fixed by the usage of context.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	go routineCounter()

	for {
		done := task()
		fmt.Println(<-done < 10)
	}
	//time.Sleep(5 * time.Second)
}

func task() chan int {
	ch := make(chan int)
	done := make(chan int)

	go func() {
		d := time.Now().Add(100 * time.Millisecond)
		ctx, cancel := context.WithDeadline(context.Background(), d)
		defer cancel()

		for i := 0; i < 10; i++ {
			go subtask(ctx, ch)
		}

		cnt := 0
		for i := 0; i < 10; i++ {
			select {
			// read from the subtasks
			case rst := <-ch:
				// if the subtask timeouts, it send 0 instead of not sending anything
				if rst > 0 {
					// counting all the successful cases(not timeout)
					cnt++
				}
			}
		}
		done <- cnt
	}()

	return done
}

func subtask(cxt context.Context, in chan int) {
	wait := rand.Intn(60) + 80
	select {
	case <-time.After(time.Duration(wait) * time.Millisecond):
		// wait for a while before sending back the result
		// this emulates some scenarios like http
		in <- wait
	case <-cxt.Done():
		in <- 0
	}
}

func routineCounter() {
	// this is used to count the number of live goroutines
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("the number of goroutines is %d\n", runtime.NumGoroutine())
	}
}
