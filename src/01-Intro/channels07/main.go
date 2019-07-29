package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// START OMIT
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)
		defer fmt.Println(&stdoutBuff, "Producer Done...")

		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received: %v\n", integer)
	}
	// END OMIT
}
