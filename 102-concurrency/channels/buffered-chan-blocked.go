package main

import "fmt"

func main() {
	bufferedChan := make(chan int, 5)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	bufferedChan <- 4
	bufferedChan <- 5
	bufferedChan <- 6 // Burada hata vermektedir. çünkü tüm goroutine'ler uyuyor

	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
}
