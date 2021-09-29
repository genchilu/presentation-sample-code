package main

import (
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func doSomething() {
	for i := 0; i < 10000; i++ {
		rand.Intn(1000000)
	}
}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	_ = trace.Start(f)
	defer trace.Stop()

	for i := 0; i < 10000; i++ {
		go doSomething()
	}
	time.Sleep(10 * time.Minute)
}
