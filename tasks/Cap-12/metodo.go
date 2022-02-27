// Método

package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) apresentacao() (string, string) {
	return "Olá, meu nome é", p.nome
}

func main() {
	pessoa1 := pessoa{"Joao", 30}

	fmt.Println(pessoa1.apresentacao())
}
