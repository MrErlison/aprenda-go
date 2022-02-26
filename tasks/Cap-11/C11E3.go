// C11E3

package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracao bool
}

type sedan struct {
	veiculo
	luxo bool
}

func main() {

	carro1 := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "preto"},
		luxo: true}

	carro2 := caminhonete{veiculo{2, "verde"}, false}

	carro3 := veiculo{4, "branca"}

	fmt.Println("sedan", carro1.cor, "é luxo?", carro1.luxo)
	fmt.Println("caminhonete", carro2.cor, "quantidade de portas", carro2.portas)
	fmt.Println("veiculo", carro3.cor)
}
