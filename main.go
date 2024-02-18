package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func sayWorld() {
	for i := 0; i < 5; i++ {
		fmt.Println("World")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Start two goroutines concurrently
	go sayHello()
	go sayWorld()

	// Allow time for goroutines to complete
	time.Sleep(1 * time.Second)
}
