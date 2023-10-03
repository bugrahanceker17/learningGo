package main

import "fmt"

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

	fmt.Println("hello channels")

	/*
		Output is non-deterministic. Scheduler probably will not have time to schedule goroutines.
		So we will not see channel value in the output
	*/

	/*
		Çıktı deterministik değildir. Zamanlayıcının muhtemelen goroutinleri planlamak için zamanı olmayacaktır.
		Yani çıktıda kanal değerini görmeyeceğiz
	*/

}
