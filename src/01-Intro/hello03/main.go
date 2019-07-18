package main

import "fmt"

//START OMIT
func main() {
	sayHello := func() {
		fmt.Println("hello")
	}

	go sayHello()
}

//END OMIT
