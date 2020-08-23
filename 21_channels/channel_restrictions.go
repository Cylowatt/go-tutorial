package main

import "fmt"

func channelRestrictions() {
	ch := make(chan int)
	wg.Add(2)

	// Data flowing out of the channel.
	// Receive-only channel.
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println("Received A:", i)
		wg.Done()
	}(ch)

	// Data flowing into the channel.
	// Send-only channel.
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)

	wg.Wait()
}
