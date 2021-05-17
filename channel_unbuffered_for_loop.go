package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {

			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 100
			wg.Done()
		}()
	}
	wg.Wait()
}
