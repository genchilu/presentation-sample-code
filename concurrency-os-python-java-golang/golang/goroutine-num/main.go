package main

import (
	"time"
)

func doSomething() {
	time.Sleep(100000000)
}

func main() {

	for i := 0; i < 10000; i++ {
		go doSomething()
	}

	time.Sleep(10 * time.Minute)

}
