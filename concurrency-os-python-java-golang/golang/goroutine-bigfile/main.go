package main

import (
	"io"
	"os"
	"runtime/trace"
	"sync"
)

var wg sync.WaitGroup

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

	wg.Done()
}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	_ = trace.Start(f)
	defer trace.Stop()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go readBigFile()
	}

	wg.Done()
}
