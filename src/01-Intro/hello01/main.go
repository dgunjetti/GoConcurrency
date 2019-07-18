package main

import "fmt"

// START OMIT
func main() {
	go sayHello()
}

func sayHello() {
	fmt.Println("hello")
}

// END OMIT
