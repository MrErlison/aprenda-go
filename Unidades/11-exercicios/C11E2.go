// C11E2

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	mapa := make(map[string]pessoa)

	mapa["pessoa1"] = pessoa{
		nome:      "nome1",
		sobrenome: "sobrenome1",
		sabores:   []string{"sabor1", "sabor2", "sabor3"}}

	for _, x := range mapa {
		fmt.Println(x.nome, x.sobrenome)
		for _, y := range x.sabores {
			fmt.Println(" *", y)
		}
	}

}
