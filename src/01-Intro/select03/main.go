package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("time out")
	}
	// END OMIT
}
