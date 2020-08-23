package main

import (
	"fmt"
	"sync"
)

var m = sync.RWMutex{}

func mutexes() {
	counter = 0

	for i := 0; i < 10; i++ {
		wg.Add(2)

		// Lock in the sync parent context; unlock asynchronously.
		// Everything runs synchronously,
		// so there goroutines are useless.
		m.RLock()
		go printCounter2()
		m.Lock()
		go increment2()
	}

	wg.Wait()
}

func printCounter2() {
	fmt.Println("Counter:", counter)
	m.RUnlock()
	wg.Done()
}

func increment2() {
	counter++
	m.Unlock()
	wg.Done()
}