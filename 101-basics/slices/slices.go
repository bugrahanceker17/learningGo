package main

import "fmt"

var slc1 []int

// array'lerin belirli bir boyutu varken, slices'lar büyüyebilir boyutlara sahiptir
func main() {
	slc2 := make([]int, 0, 3)

	slc2 = append(slc2, 1)
	slc2 = append(slc2, 24)
	slc2 = append(slc2, 5675)
	// başta verdiğimiz kapasiteyi aşan bir yeni değer eklediğimizde 3 verdiğimiz kapasiteyi 6 olarak değiştirir. yani 2 katına çıkartır
	slc2 = append(slc2, 56751231)

	fmt.Println(slc1, slc2)

	fmt.Printf("slc1 len: %d cap: %d\n", len(slc1), cap(slc1))
	fmt.Printf("slc2 len: %d cap: %d\n", len(slc2), cap(slc2))

}
