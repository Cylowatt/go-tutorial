package main

import "fmt"

// Matches the io.Writer interface.
type Writer interface {
	// Naming convention: if there is only one method, name the
	// interface after the method, adding "-er" at the end.
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(data []byte) (int, error) {
	return fmt.Println(string(data))
}

type Closer interface {
	Close() error
}

// Interface composition.
// Same idea as embedding structs.
type WriteCloser interface {
	Writer
	Closer
}