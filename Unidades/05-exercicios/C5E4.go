// C5E4
// - Crie um programa que:
//     - Atribua um valor int a uma variável
//     - Demonstre este valor em decimal, binário e hexadecimal
//     - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
//     - Demonstre esta outra variável em decimal, binário e hexadecimal

package main

import (
	"fmt"
)

func main() {
	x := 100
	fmt.Printf("%d, %#x, %b\n", x, x, x)
	x = x << 1
	fmt.Printf("%d, %#x, %b\n", x, x, x)
	x = x >> 1
	fmt.Printf("%d, %#x, %b", x, x, x)
}
