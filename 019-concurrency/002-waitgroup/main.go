package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	print("Hello, World!")
	print("Programming in Go!")
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	// ----------------------------------------------------------------------

	start = time.Now()

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		print("Hello, World!")
		waitGroup.Done() // -1
	}()

	go func() {
		print("Programming in Go!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0

	elapsed = time.Since(start)
	fmt.Println(elapsed)
}

func print(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
