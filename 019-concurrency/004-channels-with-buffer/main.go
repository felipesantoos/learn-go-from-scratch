package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Hello, World!"
	channel <- "Programming in Go!"

	// message := <-channel
	// fmt.Println(message)
	// message = <-channel
	// fmt.Println(message)

	// close(channel)
	// for message := range channel {
	// 	fmt.Println(message)
	// }

	// for i := 0; i < 2; i++ {
	// 	message := <-channel
	// 	fmt.Println(message)
	// }

	close(channel)
	for {
		message, isOpen := <-channel
		if !isOpen {
			break
		}
		fmt.Println(message)
	}
}
