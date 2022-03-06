package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	calculations := make(chan int, 45)
	results := make(chan int, 45)

	go worker(calculations, results)
	// go worker(calculations, results)

	for i := 0; i < 45; i++ {
		calculations <- i
	}
	close(calculations)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func worker(calculations <-chan int, results chan<- int) {
	for calculation := range calculations {
		results <- fibonacci(calculation)
	}
}

func fibonacci(pos int) int {
	if pos == 0 || pos == 1 {
		return pos
	} else {
		return fibonacci(pos-1) + fibonacci(pos-2)
	}
}
