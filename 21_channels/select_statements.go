package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

// struct{} has zero memory allocation.
//
// chan struct{} - signal-only channel.
// chan struct{} sends no data through, except for
// the fact that a message was sent.
func selectStatements() {
	doneCh := make(chan struct{})
	logCh := make(chan logEntry)

	go logger(logCh, doneCh)

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now().Add(15 * time.Second), logInfo, "App is shutting down"}

	doneCh <- struct{}{}
}

func logger(logCh <-chan logEntry, doneCh <-chan struct{}) {
	for {
		// Select will listen until any of the channels sends a message.
		//
		// Add default: to make the select non-blocking.
		// This means that if there are messages on the channels, these
		// will be processed. If not, the select will break out of itself.
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			// When the done channel signals, break.
			break
		}
	}
}
