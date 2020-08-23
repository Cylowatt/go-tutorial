package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Printf("Introduction to Channels\n------------\n")
	intro()

	fmt.Printf("\nIntroduction 2\n------------\n")
	intro2()

	fmt.Printf("\nRead Write\n------------\n")
	readWrite()

	fmt.Printf("\nChannel Restrictions\n------------\n")
	channelRestrictions()

	fmt.Printf("\nBuffered Channels\n------------\n")
	bufferedChannels()

	fmt.Printf("\nChannel Loops\n------------\n")
	channelLoops()

	fmt.Printf("\nManual Channel Loops\n------------\n")
	manualChannelLoops()

	fmt.Printf("\nSelect Statements\n------------\n")
	selectStatements()
}
