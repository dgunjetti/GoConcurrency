package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT1
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}
	// END OMIT1

	// START OMIT2
	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		// cancel the operation after 1 sec
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling dowork goroutine...")
		close(done)
	}()
	<-terminated
	fmt.Println("done...")
	// END OMIT2
}
