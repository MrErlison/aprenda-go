// C20E3
// - Utilizando goroutines, crie um programa incrementador:
//     - Tenha uma variável com o valor da contagem
//     - Crie um monte de goroutines, onde cada uma deve:
//         - Ler o valor do contador
//         - Salvar este valor
//         - Fazer yield da thread com runtime.Gosched()
//         - Incrementar o valor salvo
//         - Copiar o novo valor para a variável inicial
//     - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
//     - Demonstre que há uma condição de corrida utilizando a flag -race

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitgroup sync.WaitGroup

func main() {

	contador := 10
	yd := 0

	waitgroup.Add(contador)
	for i := 0; i < contador; i++ {
		fmt.Println("Goroutine #", i)
		go func() {
			v := yd
			runtime.Gosched()
			v++
			yd = v
			waitgroup.Done()
			fmt.Println("Contador #", contador)
		}()
	}
	waitgroup.Wait()

}
