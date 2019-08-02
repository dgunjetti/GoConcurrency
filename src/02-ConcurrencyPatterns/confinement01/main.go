package main

import "fmt"

func main() {
	// START OMIT1
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}
	// END OMIT1
	// START OMIT2
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("recived: %d\n", result)
		}
		fmt.Println("done receiving")
	}

	results := chanOwner()
	consumer(results)
	// END OMIT2
}
