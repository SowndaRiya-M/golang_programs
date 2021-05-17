package main

import (
	"fmt"
)

func main() {
	switch 2 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3, 4, 5:
		fmt.Println("From 3 to 5")
	default:
		fmt.Println("another number")
	}
}
