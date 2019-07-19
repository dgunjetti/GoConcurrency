package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT1
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			wg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			fn()
		}()
		wg.Wait()
	}
	// END OMIT1
	// START OMIT2
	var wg sync.WaitGroup
	wg.Add(2)
	subscribe(button.Clicked, func() {
		fmt.Println("mouse clicked")
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("maximize window")
		wg.Done()
	})
	button.Clicked.Broadcast()
	wg.Wait()
	// END OMIT2
}
