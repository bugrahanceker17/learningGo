package main

import (
	"fmt"
	"io"
	"time"
)

// bir yandan yazma işlemi yaparken bir yandan okuma işlemi yapılacağı zaman kullanılır
func main() {

	pRead, pWriter := io.Pipe()
	done := make(chan struct{})

	go read(pRead, done)
	go write(pWriter)

	<-done
}

func read(reader *io.PipeReader, done chan struct{}) {
	// empty struct kullanılmasının amacı yer kaplamaması içindir

	buff := make([]byte, 1024)
	for {
		readed, err := reader.Read(buff)
		if readed == 0 {
			if err == io.EOF {
				done <- struct{}{}
				break
			}

			if err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
		} else {
			fmt.Println(buff[:readed])

			if err == io.EOF || err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
		}
	}
}

func write(writer *io.PipeWriter) {
	i := 0
	for {
		if i == 10 {
			writer.Close()
			break
		}

		writer.Write([]byte(string(i)))
		i++
		time.Sleep(time.Millisecond * 500)
	}
}
