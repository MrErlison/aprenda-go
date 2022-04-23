// C13E4
// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.

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
