package main

import "fmt"

// START OMIT
func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "hello channels"
	}()
	msg, ok := <-stringStream
	fmt.Printf("(%v): %v\n", ok, msg)
}

// END OMIT
