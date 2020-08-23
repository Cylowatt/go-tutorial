package main

import (
	"fmt"
)

// Buffered channels are used when your senders generate data in bursts
// and you don't want to block them before your receivers can handle it.
func bufferedChannels() {
	// Buffer size = 50.
	ch := make(chan int, 50)
	wg.Add(2)

	// Receives a single message.
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// Sends two messages.
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)

	wg.Wait()
}

func channelLoops() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {

		// Constantly wait for messages available on the channel
		// until the channel is closed.
		for i := range ch {
			fmt.Println(i)
		}

		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27

		// If you try to send a message to a closed channel,
		// the application will panic.
		//
		// There is no way to check if the channel is closed
		// or to open it again.
		close(ch)

		wg.Done()
	}(ch)

	wg.Wait()
}

func manualChannelLoops() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for {
			// Manually check whether the channel is closed.
			// ok will be false if the channel is closed.
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}

		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
