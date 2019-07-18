package main

import (
	"fmt"
	"sync"
)

func main() {
	//START OMIT1
	var count int
	var lock sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		count++
	}()
	//END OMIT1
	//START OMIT2
	wg.Add(1)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		count--
	}()
	wg.Wait()
	fmt.Println("complete..")
	//END OMIT2
}
