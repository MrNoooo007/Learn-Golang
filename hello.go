package main

import "fmt"

func main() {
	i := func(a, b int) int {
		return a + b
	}
	fmt.Print(i(1, 2))
}
