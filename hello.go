package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // to read counter
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello ", counter)
	m.RUnlock() // unlock counter
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
