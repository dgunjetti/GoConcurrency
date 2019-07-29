package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	start := time.Now()
	var c1 <-chan int
	select {
	case <-c1:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
	// END OMIT
}
