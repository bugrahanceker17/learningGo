package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	str1 := strings.NewReader("read string 1\n")
	str2 := strings.NewReader("read string 2\n")

	multiReader := io.MultiReader(str1, str2)

	_, err := io.Copy(os.Stdout, multiReader)
	if err != nil {
		return
	}
}
