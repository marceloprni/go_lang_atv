package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 100)
	go setList(channel)

	valor := <-channel
	fmt.Println(valor)

	for v := range channel {
		fmt.Println(v)
		time.Sleep(time.Second)
	}

}

func setList(channel chan<- int) {
	for j := 0; j < 100; j++ {
		channel <- j
		fmt.Println("enviando :", j)
	}
	close(channel)
}
