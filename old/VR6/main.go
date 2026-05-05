package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go callApi(&wg)
	go processInternal(&wg)
	fmt.Println("All tasks completed")
	wg.Wait()

}

func callApi(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("API call completed")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Internal processing completed")
	wg.Done()
}
