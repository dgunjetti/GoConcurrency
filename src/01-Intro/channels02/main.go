package main

import "fmt"

// START OMIT
func main() {
	stringStream := make(chan string)

	go func() {
		if 0 != 1 {
			return
		}
		// block until there is reader
		stringStream <- "hello channels"
	}()
	// block until a value is placed on the channel
	fmt.Println(<-stringStream)
}

// END OMIT
