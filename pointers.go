package main

import (
	"fmt"
)

func main() {
	a := 100
	b := &a
	fmt.Println(a, b)
	fmt.Println(a, *b)
}
