package main

import (
	"io"
	"os"
	"time"
)

func readBigFile() {
	fi, err := os.Open("bigfile")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
	}
}

func main() {
	for i := 0; i < 2000; i++ {
		go readBigFile()
	}

	time.Sleep(10 * time.Minute)
}
