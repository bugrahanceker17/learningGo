package main

import "fmt"

func main() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	chan2 := make(chan int, 2)
	chan2 <- 2

	//this will not print the second chan value
	// normalde select'teki case'ler verilen sırayla gider ancak bir channel dinlendiğinde random olarak çalışmaktadır
	select {
	case c1Val := <-chan1:
		fmt.Println(c1Val)
	case c2Val := <-chan2:
		fmt.Println(c2Val)
	}

	//tüm chan'deki verileri okuyabilmek için
	var done bool
	for !done {
		select {
		case c1Val := <-chan1:
			fmt.Println(c1Val)
		case c2Val := <-chan2:
			fmt.Println(c2Val)
		default:
			break
		}
	}

}
