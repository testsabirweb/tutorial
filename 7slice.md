## Slice
```go
package main

import (
	"fmt"
	"slices"
)

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("s:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	c = append(c, []string{"p", "q", "r"}...) ///another good to know method for slice in slice
	fmt.Println("cpy:", c)

	//////////////////////////////////
	l := s[2:5]
	fmt.Println("s:", s, "l:", l)
	fmt.Println("sl1:", l)
	l[0] = "QQQQQQQQQQQQ" //change in l will also reflect in s (opposite of python)
	fmt.Println("s:", s, "l:", l)
	/////////////////////////////////

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) { // cannot use == to compare
		fmt.Println("t == t2 using slices.Equal")
	}
}
```