package main

import "fmt"

// START OMIT
func main() {
	valueStream := make(chan int)
	close(valueStream)
	integer, ok := <-valueStream
	fmt.Printf("(%v), %v\n", ok, integer)
}

// END OMIT
