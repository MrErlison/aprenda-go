// C13E10
// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.

package main

import "fmt"

func Sequencia() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	proximoInteiro := Sequencia()
	avancarInteiro := Sequencia()

	fmt.Println(proximoInteiro())
	fmt.Println(proximoInteiro())

	fmt.Println(avancarInteiro())

	fmt.Println(proximoInteiro())

	fmt.Println(avancarInteiro())

}
