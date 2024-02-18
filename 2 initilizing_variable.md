## Initilization
```go
package main

import "fmt"

func main() {
	// Using make function
	sliceMake := make([]int, 3, 5)
	fmt.Println("Slice initialized with make:", sliceMake)
	fmt.Println(cap(sliceMake), len(sliceMake))

	sliceMake = append(sliceMake, 2, 3, 4, 5)
	fmt.Println("Slice initialized with make:", sliceMake)
	fmt.Println(cap(sliceMake), len(sliceMake))

	// Using a composite literal
	fruits := []string{"apple", "orange", "banana"}

	// Using []Type{} syntax for an empty slice
	animals := []string{}

	// Initializing an empty slice and appending elements
	emptySlice := []int{}
	emptySlice = append(emptySlice, 1, 2, 3)

	// Displaying the initialized slices
	fmt.Println("Slice initialized with make:", sliceMake)
	fmt.Println("Slice initialized with composite literal:", fruits)
	fmt.Println("Empty slice initialized with []Type{} syntax:", animals)
	fmt.Println("Empty slice initialized and appended with elements:", emptySlice)
}

```

```bash
Slice initialized with make: [0 0 0]
Slice initialized with composite literal: [apple orange banana]
Empty slice initialized with []Type{} syntax: []
Empty slice initialized and appended with elements: [1 2 3]
````