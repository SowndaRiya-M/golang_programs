package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "HElllllloooooo"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Thank you"
	time.Sleep(100 * time.Microsecond)
}
