package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	var init int
	moduleInit := func() { init++ }
	var once sync.Once
	once.Do(moduleInit)
	once.Do(moduleInit) // redundant
	fmt.Printf("init=%d", init)
	// END OMIT
}
