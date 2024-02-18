## Switch
```go
package main

import "fmt"

func main() {
    day := "Wednesday"

    switch day {
    case "Monday":
        fmt.Println("It's the start of the week.")
    case "Wednesday":
        fmt.Println("It's the middle of the week.")
    case "Friday":
        fmt.Println("It's the end of the week.")
    default:
        fmt.Println("It's another day.")
    }
}
```
