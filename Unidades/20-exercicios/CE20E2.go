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
