package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := write("Channel 1", time.Second)
	channel2 := write("Channel 2", time.Second*2)

	start := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel1)
		fmt.Println(<-channel2)
	}
	fmt.Println(time.Since(start))

	start = time.Now()
	for i := 0; i < 10; i++ {
		select {
		case message := <-channel1:
			fmt.Println(message)
		case message := <-channel2:
			fmt.Println(message)
		}
	}
	fmt.Println(time.Since(start))
}

func write(text string, sleep time.Duration) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)
			time.Sleep(time.Second)
		}
	}()

	return channel
}
