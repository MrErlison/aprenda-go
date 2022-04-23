// C20E1
// - Alem da goroutine principal, crie duas outras goroutines.
// - Cada goroutine adicional devem fazer um print separado.
// - Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
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
