## rune vs byte vs string

```go
package main

import "fmt"

func main() {
    // String
    str := "Hello, World!" //string is a sequence of characters represented as a slice of bytes.// UTF-8 encoded
    fmt.Println("String:", str)

    // Rune
    runeValue := '世' //int32 , used for unicode
    fmt.Println("Rune:", runeValue)

    // Byte
    byteValue := byte('A') //uint8 , usually used for ascii 
    fmt.Println("Byte:", byteValue)
}
```