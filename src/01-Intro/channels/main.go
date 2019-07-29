package main

import "fmt"

// START OMIT
func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	fmt.Println(<-stringStream)
}

// END OMIT
