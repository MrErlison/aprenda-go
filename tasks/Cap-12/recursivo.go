// func recursivo

package main

import "fmt"

func main() {
	fmt.Println(fatorial(4))
}

func fatorial(valor int) int {

	if valor == 0 {
		return 0
	} else if valor == 1 {
		return valor
	}

	return valor * fatorial(valor-1)
}
