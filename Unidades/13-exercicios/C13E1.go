// C13E1
// - Crie uma função que retorne um int
// - Crie outra função que retorne um int e uma string
// - Chame as duas funções
// - Demonstre seus resultados

package main

import "fmt"

func retorno() int {
	return 10
}

func retornos() (int, string) {
	return 10, "dez"
}

func main() {
	fmt.Println(retorno())
	fmt.Println(retornos())
}
