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