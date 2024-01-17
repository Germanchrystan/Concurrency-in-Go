package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) { // We need to pass the input to the goroutine
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

/*
Previously the go routine didn't receive the int i as an argument. When we ran this script the output was:
4
4
4

This is because by the time the goroutine got the chance to run the value of i, it had been already incremented to four.

Goroutines operate on the current value of the variable at the time of their execution. If we want the goroutine to operate on a specific value,
then we need to pass that as an input to the goroutine.
*/
