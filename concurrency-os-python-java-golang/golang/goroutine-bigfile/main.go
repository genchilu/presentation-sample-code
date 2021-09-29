package main

import (
	"io"
	"math/rand"
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
}

func doSomething() {
	for i := 0; i < 10; i++ {
		rand.Intn(1000000)
	}
	wg.Done()
}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	_ = trace.Start(f)
	defer trace.Stop()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething()
	}

	wg.Wait()
}
