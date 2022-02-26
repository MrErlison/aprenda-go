// C11E1

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{
		nome:      "nome1",
		sobrenome: "sobrenome1",
		sabores:   []string{"sabor1", "sabor2", "sabor3"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, x := range pessoa1.sabores {
		fmt.Println("\t", x)
	}

	pessoa2 := pessoa{"nome2", "sobrenome2",
		[]string{"sabor1", "sabor2", "sabor3"}}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, x := range pessoa2.sabores {
		fmt.Println("\t", x)
	}
}
