package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//START OMIT1
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func() {
		time.Sleep(1 * time.Second)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}
	// END OMIT1
	// START OMIT2
	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue()
		c.L.Unlock()
	}
	// END OMIT2
}
