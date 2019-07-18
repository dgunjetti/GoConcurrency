package main

import "fmt"

//START OMIT
func main() {
	go func() {
		fmt.Println("hello")
	}()
}

//END OMIT
