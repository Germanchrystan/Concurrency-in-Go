package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			// Send iterator over channel
			ch <- i
		}
		close(ch)
	}()

	// Range over channel to receive values
	for v := range ch {
		fmt.Println(v)
	}
}