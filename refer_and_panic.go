package main

import "fmt"

func recovery() {
	fmt.Println("This is recovery function...")
}

func Panic() {
	defer recovery()
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}

func main() {
	Panic()
	fmt.Println("Main block is executed completely...")
}
