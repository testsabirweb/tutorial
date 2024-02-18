package main

import "fmt"

func main() {
    defer fmt.Println("First deferred statement.")
    defer fmt.Println("Second deferred statement.")
    defer fmt.Println("Third deferred statement.")

    fmt.Println("Regular execution.")
}
