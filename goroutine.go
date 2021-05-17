package main

import (
	"fmt"
	"time"
)

func main() {
	go gorout()
	time.Sleep(100 * time.Microsecond)
}

func gorout() {
	fmt.Println("GO ROUTINE PROGRAM...OH..YEAHHHHHHHH")
}
