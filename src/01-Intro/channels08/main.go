package main

import "fmt"

func main() {
	// START OMIT
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 6)
		go func() {
			defer close(resultStream)
			for i := 0; i < 6; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d", result)
	}
	//END OMIT
	fmt.Println("Done receiving")

}
