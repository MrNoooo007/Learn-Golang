package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	ch1 := make(chan int, 50) // create a channel with 50 buffers
	ch2 := make(chan int, 50) // create a channel with 50 buffers

	wg.Add(2)
	go func() {
		for {
			select {
			case i := <-ch1:
				fmt.Print("Channel 1 ", i)
			case j := <-ch2:
				fmt.Print("channel 2 ", j)
			default:
				break
			}
		}
		wg.Done()
	}()
	go func() {
		ch1 <- 42  // Set value to ch channel
		close(ch1) // close channel

		ch2 <- 50
		close(ch2)
		wg.Done()
	}()
	wg.Wait()
}
