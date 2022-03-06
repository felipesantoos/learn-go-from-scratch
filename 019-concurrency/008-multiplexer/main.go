package main

import (
	"fmt"
	"time"
)

func main() {
	channel := multiplex(
		write("Channel 1", time.Second/2),
		write("Channel 2", time.Second*2),
	)
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplex(inputChannel1, inputChannel2 <-chan string) <-chan string {
	outputChannel := make(chan string)
	go func() {
		for {
			select {
			case message := <-inputChannel1:
				outputChannel <- message
			case message := <-inputChannel2:
				outputChannel <- message
			}
		}
	}()
	return outputChannel
}

func write(text string, sleep time.Duration) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("Receveid value: %s", text)
			time.Sleep(sleep)
		}
	}()
	return channel
}
