package main

import (
	"fmt"
	"time"
)

/*
 - Concurrency != paralelism.

 - Paralelism:
   - To execute one or more tasks at the same time.
   - It's possible with CPUs with more than one nucleus.

 - Concurrency
   - Tasks can be executed in parallel (if the CPU have more than one nucleus).
   - Can be used with CPUs with only one nucleus.
   - One task don't need to wait other task finish to be executed.
   - T1 → T2 → T1 → T2... (alternately)
*/

func main() {
	go print("Hello, World!")
	print("Programming in Go!")
}

func print(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
