package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for j := 0; j <= 2; j++ {
			c1 <- "from 1"
		}
	}()

	go func() {
		for i := 0; i <= 2; i++ {
			c2 <- "from 2"
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
				//default:
				//	fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
