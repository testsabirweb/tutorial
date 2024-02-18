## Different ways to sort

```go
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	arr := []int{5, 3, 2, 7, 4, 6}
	sort.Ints(arr)
	fmt.Println(arr)

	arr2 := []string{"sjfh", "eihrkfn", "sj"}
	sort.Strings(arr2)
	fmt.Println(arr2)

	runeSlice := []rune{'a', 'z', 'g', 'c', 'b'}
	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})
	fmt.Println(string(runeSlice))

	arr3 := make([]Person, 0)
	arr3 = append(arr3, Person{Name: "sabir", Age: 26}, Person{"something", 20}, Person{"anythong", 21})
	sort.Slice(arr3, func(i, j int) bool {
		return arr3[i].Age < arr3[j].Age
	})
	fmt.Println(arr3)

	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(ByAge(family))
	fmt.Println(family)

}```
