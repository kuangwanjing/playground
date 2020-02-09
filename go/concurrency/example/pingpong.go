package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)

	go player(context.Background(), "bob", table)
	go player(context.Background(), "alice", table)
	go player(context.Background(), "kate", table)

	table <- Ball
	time.Sleep(5 * time.Second)
}

func player(cxt context.Context, player string, table chan int) {
	for {
		select {
		case ball := <-table:
			ball++
			fmt.Printf("player %s gets the ball %d\n", player, ball)
			time.Sleep(100 * time.Millisecond)
			table <- ball
		case <-cxt.Done():
			break
		}
	}
}
