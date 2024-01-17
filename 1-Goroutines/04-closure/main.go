package main

import (
	"fmt"
	"sync"
)

//TODO: run the program and check that variable i
// was pinned for access from goroutine even after
// enclosing function returns.

func main() {
	var wg sync.WaitGroup

	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("value of i: %v\n", i)
		}()
		fmt.Println("return from function")
		return
	}

	incr(&wg)
	wg.Wait()
	fmt.Println("done..")
}

/*
OUTPUT:
return from function
value of i: 1
done...

We have a function, inside we have a local variable.
We are spinning a goroutine and we are returning from the function.
Inside the goroutine, we are accessing the local variable of the function,
and we are incrementing its value and we are printing the value.

In the main routine we are calling a function and we are waiting for the goroutines to execute.

The function has returned, but goroutine still has the access to the local
variable of the function.

So usually when the function returns, the local variables goes out of scope.
But here, the runtime is clever enough to see that the reference to a local variable i is still
being held by the goroutine,

so it pins it, it moves it from the stack to heap, so that goroutine still has the access to the variable
even after the enclosing function returns.
*/
