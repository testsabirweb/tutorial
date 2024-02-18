## waitgroup
**sync.WaitGroup** is used to wait for a collection of goroutines to finish before proceeding.
It ensures that the main goroutine does not exit before other goroutines are done.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine is done
	fmt.Printf("Worker %d starting\n", id)
	// Simulate some work
	time.Sleep(time.Second * 1)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All workers done")
}
```
