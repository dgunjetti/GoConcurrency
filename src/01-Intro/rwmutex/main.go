package main

import "sync"

func main() {
	// START OMIT1
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
		//compute
	}
	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
		// read computed value
	}
	// END OMIT1
	// START OMIT2
	var wg sync.WaitGroup
	var m sync.RWMutex

	wg.Add(1)
	go producer(&wg, &m)
	go observer(&wg, m.RLocker())
	//END OMIT2
}
