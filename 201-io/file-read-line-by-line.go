package main

import (
	"bufio"
	"fmt"
	"os"
)

var path = "./201-io/lines.txt"

func main() {
	err := os.WriteFile(path, []byte("line1\nline2\nline3\nline4"), os.ModePerm)
	if err != nil {
		return
	}

	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
