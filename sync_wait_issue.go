package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

//var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go Hello()
		go increment()
	}
	wg.Wait()
}
func Hello() {
	//m.RLock()
	fmt.Printf("Hello %v\n", counter)
	//m.RUnlock()
	wg.Done()
}

func increment() {
	//m.Lock()
	counter++
	//m.Unlock()
	wg.Done()
}
