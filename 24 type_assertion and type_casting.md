## Type Assertion
Type assertion is used to extract the underlying concrete value from an interface type.

```go
package main

import "fmt"

func main() {
	var val interface{} = 42

	// Type assertion
	if num, ok := val.(int); ok {
		fmt.Println("The value is an integer:", num)
	} else {
		fmt.Println("The value is not an integer.")
	}
}
```
*In Go, there is no direct concept of "type casting" as you might find in languages like C or C++. Instead, Go uses type assertions and explicit type conversions to work with different types.*

## Explicit Type Conversion
```go
package main

import "fmt"

func main() {
	var x int = 42
	var y float64 = float64(x)

	fmt.Println(y)
}
```
