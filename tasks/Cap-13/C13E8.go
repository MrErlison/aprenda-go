// C13E8

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
