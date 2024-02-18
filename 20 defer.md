## defer
**defer** statement is used to ensure that a function call is performed later in a program's execution, usually for purposes like cleanup or resource release. The deferred call will be executed just before the surrounding function returns.

```go
package main

import "fmt"

func main() {
    defer fmt.Println("First deferred statement.")
    defer fmt.Println("Second deferred statement.")
    defer fmt.Println("Third deferred statement.")

    fmt.Println("Regular execution.")
}
```
output:
```
Regular execution.
Third deferred statement.
Second deferred statement.
First deferred statement.
```