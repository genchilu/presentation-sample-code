package main

import (
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
)

var wg sync.WaitGroup

func doSomething() {
	for i := 0; i < 10000; i++ {
		rand.Intn(1000000)
	}
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
