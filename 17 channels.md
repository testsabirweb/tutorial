## channels
Channels facilitate communication and synchronization between goroutines.
They provide a way for one goroutine to send data to another goroutine.
Channels can be used to prevent race conditions and ensure safe communication between concurrent tasks.
```go
package main

import "fmt"

func main() {
    // Creating a channel
    ch := make(chan int)

    // Starting a goroutine to send data to the channel
    go func() {
        ch <- 42
    }()

    // Receiving data from the channel
    value := <-ch
    fmt.Println("Received:", value)
}
```
