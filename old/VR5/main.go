package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 5999; i++ {
		go showMessage(strconv.Itoa(i))
	}
	time.Sleep(time.Duration(time.Hour.Seconds() * float64(5)))
}

func showMessage(message string) { fmt.Println(message) }
