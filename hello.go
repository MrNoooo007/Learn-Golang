package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	ch := make(chan int, 50) // create a channel with 50 buffers

	wg.Add(2)
	go func(ch <-chan int) {
		// i := <-ch // get a value from channel. Waiting for value assigned in channel
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42 // Set value to ch channel
		ch <- 50
		ch <- 50
		ch <- 50
		close(ch) // close channel
		wg.Done()
	}(ch)
	wg.Wait()
}
