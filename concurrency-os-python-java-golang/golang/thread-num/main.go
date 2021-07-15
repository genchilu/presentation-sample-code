package main

import (
	"math/rand"
	"time"
)

func doSomething() {
	for {
		rand.Intn(1000000)
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		go doSomething()
	}
	time.Sleep(10 * time.Minute)
}
