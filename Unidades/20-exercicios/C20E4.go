// C20E4
// - Utilize mutex para consertar a condição de corrida do exercício anterior.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitgroup sync.WaitGroup
var mu sync.Mutex

func main() {

	contador := 10
	yd := 0

	waitgroup.Add(contador)
	for i := 0; i < contador; i++ {
		fmt.Println("Goroutine #", i)
		go func() {
			mu.Lock()
			v := yd
			runtime.Gosched()
			v++
			yd = v
			mu.Unlock()
			waitgroup.Done()
			fmt.Println("Contador #", contador)
		}()
	}
	waitgroup.Wait()

}
