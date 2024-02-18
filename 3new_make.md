## Diffrence between new and make
```go
package main

import "fmt"

func main() {
	// Using new for an integer
	intValue := new(int)
	*intValue = 42
	fmt.Println("Value created with new:", *intValue)

	// Using make for a slice
	mySlice := make([]int, 5, 10)
	fmt.Println("Slice created with make:", mySlice)
}

```
* Use new for allocating memory for value types and obtaining a `pointer` to the zeroed value.
* Use make for initializing slices, maps, and channels, which require non-zeroed initial values.

```bash
Value created with new: 42
Slice created with make: [0 0 0 0 0]
````