## if,else, else if

```go
package main

import "fmt"

func main() {
    num := 10

    if num > 15 {
        fmt.Println("Number is greater than 15")
    } else if num > 5 {
        fmt.Println("Number is between 6 and 15")
    } else {
        fmt.Println("Number is 5 or less")
    }
}
```