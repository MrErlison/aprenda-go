// C13E1

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
