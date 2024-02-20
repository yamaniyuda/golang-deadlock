package main

import "fmt"

func main() {
	message := make(chan string)

	// Main goroutine reads messages from channels,
	// but none of the other goroutines send
	// message to this channel
	fmt.Println(<-message)
}
