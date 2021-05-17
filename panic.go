package main

import "fmt"

func Panic() {
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}

func main() {
	Panic()
	fmt.Println("Main block is executed completely...")
}
