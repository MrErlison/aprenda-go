// - Crie um tipo para um struct chamado "pessoa"
// - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
// - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
// - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
// exemplo https://goporexemplos.github.io/interfaces.html

package main

import "fmt"

type pessoa struct {
	nome string
}

type humanos interface {
	falar()
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "~ Olá!")
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	john := pessoa{"John Doe"}

	john.falar()

	dizerAlgumaCoisa(&john)
}
