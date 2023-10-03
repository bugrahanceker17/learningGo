package main

import "fmt"

func main() {
	//unbuffered chan
	c1 := make(chan int)

	//writes are blocking operations
	c1 <- 1

	fmt.Println("hello concurrency")
}
