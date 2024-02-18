## Select
**select** statement is used for communication and synchronization between goroutines. It allows a goroutine to wait on multiple communication operations and proceed when one of them is ready. The select statement is often used with channels.
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from goroutine 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from goroutine 2"
	}()

	// The select statement waits for one of the channels to be ready
	// and then executes the corresponding case.
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timed out after 3 seconds.")
			return
		}

		time.Sleep(1 * time.Second)
	}
}
```