package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	go setList(channel)

	valor := <-channel
	fmt.Println(valor)

	for v := range channel {
		fmt.Println(v)
	}

}

func setList(channel chan int) {
	for j := 0; j < 1000; j++ {
		channel <- j
	}
	close(channel)
}
