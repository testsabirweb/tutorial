## Generics (go1.18)

```go
// all this fucntions are already defined in utils/utils.go
package utils

import (
	"fmt"
	"math"
	"os"
)

type NumberTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type ListTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Make2dArray[T any](numRows, numCols int) [][]T {
	matrix := make([][]T, numRows)
	buf := make([]T, numRows*numCols)
	for i, startRow := 0, 0; i < numRows; i, startRow = i+1, startRow+numCols {
		endRow := startRow + numCols
		matrix[i] = buf[startRow:endRow:endRow]
	}
	return matrix
}

func Populate2dArray[T any](matrix [][]T, value T) [][]T {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = value
		}
	}
	return matrix
}

func Make3DArray[T any](numRows, numCols, numHeight int) [][][]T {
	matrix := make([][][]T, numRows)
	buf := make([]T, numRows*numCols*numHeight)

	for i := range matrix {
		matrix[i] = make([][]T, numCols)
		for j := range matrix[i] {
			matrix[i][j] = buf[:numHeight:numHeight]
			buf = buf[numHeight:]
		}
	}
	return matrix
}

func Print2D[T NumberTypes](arr [][]T, n int, m int) {
	if n == 0 {
		n = len(arr)
	}
	if m == 0 {
		m = len(arr[0])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

func Swap[T NumberTypes](arr []T, index1, index2 int) {
	arr[index1], arr[index2] = arr[index2], arr[index1]
}

func Max[T NumberTypes](values ...T) T {
	if len(values) == 0 {
		panic("no values provided")
	}

	maxVal := values[0]
	for _, val := range values {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func Min[T NumberTypes](values ...T) T {
	if len(values) == 0 {
		panic("no values provided")
	}

	minVal := values[0]
	for _, val := range values {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func Diff[T NumberTypes](a, b T) T {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

// func genericWithStruct() { // just a example how we can use generic with struct
// 	type node[T utils.ListTypes] struct {
// 		data T
// 		prev *node[T]
// 		next *node[T]
// 	}
// }
```