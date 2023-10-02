package main

import "fmt"

var arr1 [3]int
var arr2 = [5]int{1, 2, 3, 4, 56}

func main() {
	arr3 := make([]int, 3)

	arr3[0] = 132
	arr3[1] = 5655

	fmt.Println(arr1, arr2, arr3)
	fmt.Printf("arr1 len: %d, \n", len(arr1))
	fmt.Printf("arr2 len: %d, \n", len(arr2))
}
