package main

import (
	"fmt"
	"time"
)

func main() {
	slc1 := []int{1, 2, 3}
	slc2 := []int{}

	for i, value := range slc1 {
		fmt.Printf("index: %d, value: %d\n", i, value)
	}

	for i := 0; i < 10; i++ {
		slc2 = append(slc2, i)
	}

	for _, value2 := range slc2 {
		for i := range slc1 {
			slc1[i] += value2
		}
	}

	//sonsuz döngü için

	c := time.After(5 * time.Second)
	for {
		b := false
		select {
		case <-c:
			b = true
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)

		}

		if b {
			break
		}
	}

	fmt.Println(slc1)
	fmt.Println(slc2)

}
