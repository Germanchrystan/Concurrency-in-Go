package main

func main() {
	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)
	// get the value computed from goroutine.
	// We want the result of the computation in our main routine without sharing of memory.
	r := <- ch
	fmt.Printf("computed value %v\n", r)
}
