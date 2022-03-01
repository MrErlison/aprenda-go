// C13E9

package main

import "fmt"

func funcaodefuncao() {
	fmt.Println("função que é passada retransmitida")
}

func funcao(x func()) {
	fmt.Println("Função que recebe outra como argumento")
	x()
}

func main() {
	funcao(funcaodefuncao)
}
