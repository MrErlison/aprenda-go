package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func main() {

	waitgroup.Add(2)

	go func() {
		fmt.Println("Goroutine #1")
		waitgroup.Done()
	}()

	go func() {
		fmt.Println("Goroutine #2")
		waitgroup.Done()
	}()

	waitgroup.Wait()
}
