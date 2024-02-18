package main

import (
	"fmt"
	"github.com/google/uuid"
)
// go get github.com/google/uuid
func main() {
	fmt.Println(uuid.New().String())
}