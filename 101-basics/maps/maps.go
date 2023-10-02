package main

import "fmt"

var mp1 map[int]string

func main() {
	mp1 = make(map[int]string)
	mp2 := make(map[int]int)
	mp3 := make(map[int]int, 3)

	mp1[0] = "test"
	mp2[0] = 2
	mp3[0] = 17
	mp3[1] = 14
	mp3[2] = 21

	fmt.Println(mp1, mp2, mp3)
	fmt.Printf("map1 len: %d,\n", len(mp1))
	fmt.Printf("map2 len: %d,\n", len(mp2))
	fmt.Printf("map3 len: %d,\n", len(mp3))
}
