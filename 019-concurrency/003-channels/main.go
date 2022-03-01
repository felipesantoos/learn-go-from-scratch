package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go print("Hello, World!", channel)
	fmt.Println("Start!")

	// for {
	// 	text, isOpen := <-channel
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println(text)
	// }

	for text := range channel {
		fmt.Println(text)
	}

	fmt.Println("End!")
}

func print(text string, channel chan string) {
	for i := 1; i <= 5; i++ {
		newText := fmt.Sprintf("%s %d", text, i)
		channel <- newText
		time.Sleep(time.Second)
	}

	close(channel)
}
