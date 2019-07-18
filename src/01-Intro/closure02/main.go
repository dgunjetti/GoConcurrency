package main

import (
	"fmt"
	"sync"
)

func main() {
	//START OMIT
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	//END OMIT
}
