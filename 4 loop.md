## Loops
```go
package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("#########################################################")
	// go1.22.0 edition //range over integers
	for i := range 10 {
		fmt.Println(i)
	}
	fmt.Println("#########################################################")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	fmt.Println("#########################################################")
	for index, value := range []int{3, 4, 5, 2} {
		fmt.Println("range", index, value)
	}
	fmt.Println("#########################################################")
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("#########################################################")
	for n := range []int{1, 2, 3, 4, 5} {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}


```
.
```bash
1
2
3
#########################################################
0
1
2
#########################################################
range 0 3
range 1 4
range 2 5
range 3 2
#########################################################
loop
#########################################################
1
3
````