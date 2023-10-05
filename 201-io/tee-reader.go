package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//örneğin bir http body geldi, body'i okumak, bir yere kaydetmek ve tekrar üzerinde işlem yapmak istediğimizde kullanılır
	sReader := strings.NewReader("test data")

	//duplicates the stream
	tReader := io.TeeReader(sReader, os.Stdout)

	fmt.Println("we will start")

	readedBytes, _ := io.ReadAll(tReader)

	fmt.Println("\n --- readed string ---")
	fmt.Println(string(readedBytes))
}
