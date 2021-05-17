package main

import "fmt"

func returnMessage() {
	fmt.Println("This is Simple Defer Function Execution")
}

func main() {
	defer returnMessage()
	fmt.Println("One")
	fmt.Println("Two")
	fmt.Println("Three")
}
