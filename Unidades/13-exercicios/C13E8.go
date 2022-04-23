// C13E8
// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.

package main

import "fmt"

func funcao() func(x string) {
	return func(x string) {
		fmt.Println(x)
	}
}

func main() {
	x := funcao()

	x("função que retorna uma função")
}
