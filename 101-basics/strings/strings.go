package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	str1 := str[:5]          // ilk 5 karakter
	str2 := str[len(str)-5:] // son 5 karakter

	fmt.Sprintf("str is : %s \n", str)   // output : Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	fmt.Sprintf("str1 is : %s \n", str1) // output : Lorem
	fmt.Sprintf("str2 is : %s \n", str2) // output : elit.

	if strings.EqualFold(str1, "loRem") { // true --- büyük küçük harf duyarlılığı yoktur
		fmt.Println("str1 equal to lorem")
	}

	if strings.HasPrefix(str, "Lorem") { // true --- büyük küçük harf duyarlılığı vardır
		fmt.Println("str has to 'Lorem'")
	}

}
