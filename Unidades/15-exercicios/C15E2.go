// C15E2
// - Crie um struct "pessoa"
// - Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
}

func mudeMe(p *pessoa) {
	(*p).nome = "nome alterado"
	(*p).sobrenome = "sobrenome alterado"

}

func main() {
	p := pessoa{nome: "nome", sobrenome: "sobrenome"}
	mudeMe(&p)
	fmt.Println(p.nome, p.sobrenome)
}
