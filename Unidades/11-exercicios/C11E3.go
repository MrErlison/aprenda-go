// C11E3
// - Crie um novo tipo: veículo
//     - O tipo subjacente deve ser struct
//     - Deve conter os campos: portas, cor
// - Crie dois novos tipos: caminhonete e sedan
//     - Os tipos subjacentes devem ser struct
//     - Ambos devem conter "veículo" como struct embutido
//     - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
//     - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
// - Usando os structs veículo, caminhonete e sedan:
//     - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
//     - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
// - Demonstre estes valores.
// - Demonstre um único campo de cada um dos dois.

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
