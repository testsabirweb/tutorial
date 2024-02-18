## mutex
**sync.Mutex** provides mutual exclusion, preventing data races when multiple goroutines access shared data concurrently. It is used to protect critical sections of code by ensuring that only one goroutine can access the protected resource at a time.

```go
package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    value int
    mu    sync.Mutex
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    var wg sync.WaitGroup
    counter := Counter{}

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }

    // Wait for all goroutines to finish
    wg.Wait()

    // Print the final counter value
    fmt.Println("Final Counter Value:", counter.Value())
}
```
