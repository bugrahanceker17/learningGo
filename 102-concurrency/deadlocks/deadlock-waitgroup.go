package main

import (
	"fmt"
	"sync"
)

func main() {
	//waitgroup for synchronization
	wg := sync.WaitGroup{}
	//wrong number of goroutine to wait
	// yanlış verilen bir sayıda deadlock olur
	wg.Add(2)

	go func() {
		fmt.Println("hello from go routine")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("hello concurrency")
}
