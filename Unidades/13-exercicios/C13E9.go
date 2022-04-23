// C13E9
// - Callback: passe uma função como argumento a outra função.

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
