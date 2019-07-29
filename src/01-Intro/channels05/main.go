package main

import "fmt"

// START OMIT
func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Printf("%v", integer)
	}
}

// END OMIT
