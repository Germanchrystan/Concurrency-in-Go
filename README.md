# Concurrency in Go
## Overview of Concurrency
Concurrency is about multiple things happenning at the same time in random order. Go provides a built in support for concurrency. It is not using an external library as other languages do, but higher level abstraction is built in, which makes writing concurrent code clean and elegant.

Lets take an example

~~~go
func Add(numbers []int) int64 {
    var sum int64
    for _, n := range numbers {
        sum += int64(n)
    }
    return sum
}
~~~

This is a simple addition function. It takes a slice of integers as input. It loops over those integers, computes the sum and returns the sum. If we passed a slice of millions of numbers as input, then it is going to take some time to compute the sum.

- When Add() is executed it runs on a single core.
- To make the computation run faster, we could divide the input slice into multiple slices, and run the add function on each part of the slice, in parallel on different cores. This way, we will be doing our computation much faster.

~~~go
func AddConcurrent(numbers []int) int64 {
    // utilize all cores on machine
    numOfCores := runtime.NumCpu()
    runtime.GOMAXPROCS(numOfCores)

    var sum int64
    max := len(numbers)

    sizeOfParts := max / numOfCores

    var wg sync.WaitGroup

    for i := 0; i < numOfCores; i++ {
        // Divide the input into parts
        start := i * sizeOfParts
        end := start + sizeOfParts
        part := numbers[start:end]

        // Run computation for each part in separate goroutines
        wg.Add(1)
        go func(nums []int){
            defer wg.Done()
            var partSum int64

            // Calculate sum for each part
            for _, n := range.nums {
                partSum += int64(n)
            }

            // Add sum of each part to cummulative sum
            atomic.AddInt64(&sum, partSum)
        }(part)
    }
    wg Wait()
    return sum
}
~~~

Concurrency is composition of independent execution computations, which may or may not run in parallel.

