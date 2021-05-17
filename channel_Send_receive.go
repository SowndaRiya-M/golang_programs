package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()
	go func() {
		ch <- 100
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}
