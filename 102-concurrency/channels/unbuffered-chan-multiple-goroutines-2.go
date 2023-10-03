package main

import (
	"fmt"
	"time"
)

func main() {
	//chan init
	unbufferedChan := make(chan int)

	//reader goroutine
	go func(unBufCan chan int) {
		value := <-unBufCan
		fmt.Println(value)
	}(unbufferedChan)

	//writer goroutine
	go func(unBufChan chan int) {
		unBufChan <- 1
	}(unbufferedChan)

	time.Sleep(time.Second)
	fmt.Println("hello channels")
}
