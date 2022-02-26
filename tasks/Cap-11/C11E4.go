// C11E4

package main

import "fmt"

func main() {

	veiculo := struct {
		marca     string
		portas    int
		cor       string
		sinistros map[string]string
		multa     []string
	}{
		marca:  "Honda",
		portas: 4,
		cor:    "branco",
		sinistros: map[string]string{
			"sinistro1": "pneu furado",
			"sinistro2": "falta de combustível",
		},
		multa: []string{"velociade superior a permitida", "passar no sinal vermelho"},
	}

	fmt.Println(veiculo.marca, veiculo.cor)
	fmt.Println("-", veiculo.multa[0])
	fmt.Println("-", veiculo.multa[1])
	fmt.Println(veiculo.sinistros["sinistro1"])

}
