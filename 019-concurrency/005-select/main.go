package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		i := 1
		for {
			channel1 <- fmt.Sprintf("Channel 1: [%d]", i)
			i++
			time.Sleep(time.Second / 2)
		}
	}()

	go func() {
		i := 1
		for {
			channel2 <- fmt.Sprintf("Channel 2: [%d]", i)
			i++
			time.Sleep(time.Second * 2)
		}
	}()

	start := time.Now()
	i := 0
	for message := range channel1 {
		fmt.Println(message)
		i++
		if i == 5 {
			break
		}
	}
	elapsed1 := time.Since(start)
	fmt.Println(elapsed1)

	start = time.Now()
	i = 0
	for message := range channel2 {
		fmt.Println(message)
		i++
		if i == 5 {
			break
		}
	}
	elapsed2 := time.Since(start)
	fmt.Println(elapsed2)

	start = time.Now()
	for i := 0; i < 5; i++ {
		messageChannel1 := <-channel1
		fmt.Println(messageChannel1)
		messageChannel2 := <-channel2
		fmt.Println(messageChannel2)
	}
	elapsed3 := time.Since(start)
	fmt.Println(elapsed3)

	fmt.Println(float64(elapsed1)+float64(elapsed2) == float64(elapsed3))

	start = time.Now()
	for i := 0; i < 10; i++ {
		select {
		case messageChannel1 := <-channel1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-channel2:
			fmt.Println(messageChannel2)
		}
	}
	elapsed4 := time.Since(start)
	fmt.Println(elapsed4)
}
