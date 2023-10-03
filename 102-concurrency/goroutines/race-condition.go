package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//raceCondition()
	//raceConditionFixed()
	raceConditionWithAtomic()
}

func raceCondition() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	//Shared value
	val := 0

	go func() {
		for i := 0; i < 1000000; i++ {
			val++
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			val++
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}

func raceConditionFixed() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	lock := sync.Mutex{}

	val := 0

	go func() {
		for i := 0; i < 1000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}

		wg.Done()
	}()

	wg.Wait()
	fmt.Println(val)
}

func raceConditionWithAtomic() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var val int32 = 0

	go func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1)
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1)
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}
