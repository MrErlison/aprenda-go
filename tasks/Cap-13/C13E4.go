// C13E4

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) imprimir() {
	fmt.Println(p.nome, p.sobrenome, p.idade)
}

func main() {

	p := pessoa{nome: "João", sobrenome: "da Silva", idade: 18}

	p.imprimir()

}
