package main

import (
	"os"
	"runtime/trace"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doSomething() {
	time.Sleep(100000000)
	wg.Done()
}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	_ = trace.Start(f)
	defer trace.Stop()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go doSomething()
	}

	wg.Wait()
}
