package main

import (
	"fmt"
	"io"
	"os"
)

var path1 = "./201-io/multi-writer-example1.txt"
var path2 = "./201-io/multi-writer-example2.txt"

func main() {
	txtFile1, _ := os.Create(path1)
	txtFile2, _ := os.Create(path2)

	multiWriter := io.MultiWriter(os.Stdout, txtFile1, txtFile2)

	n, err := multiWriter.Write([]byte("multi writer example"))

	txtFile1.Close()
	txtFile2.Close()

	fmt.Println(err)
	fmt.Println(n)

}
