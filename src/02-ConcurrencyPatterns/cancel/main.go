package main

import (
	"fmt"
	"time"
)

func main() {
	// START1 OMIT
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
	// END1 OMIT

	// START2 OMIT
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
	// END2 OMIT
}
