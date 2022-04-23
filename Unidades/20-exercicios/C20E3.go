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

	waitgroup.Add(contador)
	for i := 0; i < contador; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			waitgroup.Done()
			fmt.Println("Goroutine #", i)
			fmt.Println("Contador #", contador)
		}()
	}
	waitgroup.Wait()

}
