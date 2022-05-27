// C20E5
// - Utilize atomic para consertar a condição de corrida do exercício #3.

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var waitgroup sync.WaitGroup

func main() {

	var contador int32 = 10
	var j int = 10

	waitgroup.Add(j)
	for i := 0; i < j; i++ {
		//fmt.Println("Goroutine #", i)
		go func() {
			atomic.AddInt32(&contador, 1)
			atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println("Contador #", contador)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()

}
