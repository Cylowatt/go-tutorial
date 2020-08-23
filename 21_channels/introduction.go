package main

import "fmt"

func intro() {
	// Similar to an Rx Observable.
	ch := make(chan int)
	wg.Add(2)

	go func() {
		// Read i from the channel. Blocks until received.
		// The arrow is pointing in the direction of the data to flow to.
		// Much like synchronous Rx subscribe.
		i := <-ch
		fmt.Println("Received from the channel:", i)
		wg.Done()
	}()

	go func() {
		data := 42
		fmt.Println("Sending to the channel:", data)

		// Send the number to the channel synchronously.
		// By default, there is no buffering in the channel (1 msg at a time),
		// so this may block until the previous int has been processed.
		ch <- data
		wg.Done()
	}()

	wg.Wait()
}

func intro2() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(2)

		go func() {
			fmt.Println("Spawn Receiver")
			j := <-ch
			fmt.Println("Received:", j)
			wg.Done()
		}()

		go func(j int) {
			fmt.Println("Spawn Sender")
			ch <- 42 + j
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func readWrite() {
	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println("Received A:", i)
		ch <- 27
		wg.Done()
	}()

	go func() {
		ch <- 42
		fmt.Println("Received B:", <-ch)
		wg.Done()
	}()

	wg.Wait()
}
