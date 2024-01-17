# Go Scheduler
Go scheduler is part of the Go runtime. It is built into the executable of the application. It is known as M:N scheduler.

Goroutines are scheduled on the OS threads by the Go scheduler. Go runtime creates a number of worker OS threads, equal to ***GOMAXPROCS*** enviroment variable value. The default value is the number of processors on the machine. So if we have 4 cores, then 4 OS threads will be created.

It is the responsibility of the Go Scheduler to distribute runnable goroutines over multiple OS threads that are created. At any time, ***N*** goroutines coul be scheduled on ***M*** OS threads that runs on at most ***GOMAXPROCS*** numbers of processors.

As of GO 1.14, the Go scheduler implements asynchronous preemption. This prevents long running goroutines from hogging onto CPU. that could block other goroutines.
In asynchronous preemption, what happens is that a goroutine is given a time slice of ten milliseconds for execution. When that time slice is over, Go scheduler will try to preempt it. This provices other goroutines the opportunity to run even when there are long running CPU bound goroutines scheduled. 

Similar to threads, goroutines also have states.

- When it is created, it will be in runnable state, waiting in the run queue.
- It moves to the Executing state once the goroutine is cheduled on the OS thread.
If the goroutine runs through its time twice, then it is preempted and placed back into the run queue.
- If the goroutine gets blocked on any condition. like blocked on channel, block on a syscall or waiting for the mutex lock, then they are moved to waiting state.
Once the I/O operation is complete, they are moved back to the runnable state.


