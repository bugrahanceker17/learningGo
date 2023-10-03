package main

import "fmt"

func main() {
	//channel initialization
	unbufferedChan := make(chan int)

	//channel declaration
	var unbufferedChan2 chan int

	fmt.Println(unbufferedChan)
	fmt.Println(unbufferedChan2) // output : nil

	//only can read from chan
	go func(unbufChan <-chan int) {
		// buraya veri gelene kadar goroutine bloklanır
		value := <-unbufChan
		fmt.Println(value)

		//unbufChan <- 5 this is not work here
	}(unbufferedChan)

	unbufferedChan <- 1
}
