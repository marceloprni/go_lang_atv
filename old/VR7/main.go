package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	i := 0
	for j := 0; j < 1000; j++ {
		go func() {
			m.Lock()
			i++
			m.Unlock()
		}()
	}

	time.Sleep(time.Second * 5)
	fmt.Println(i)

}
