// C13E2
// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.

package main

import "fmt"

func soma(v ...int) int {
	total := 0
	for _, x := range v {
		total += x
	}
	return total
}

func soma_si(v []int) int {
	total := 0
	for _, x := range v {
		total += x
	}
	return total
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(soma(x...))
	fmt.Println(soma_si(x))
}
