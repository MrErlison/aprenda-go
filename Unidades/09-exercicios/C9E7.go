// C9E7
// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//     - "Nome", "Sobrenome", "Hobby favorito"
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import "fmt"

func main() {

	x := [][]string{
		[]string{"nome1", "sobrenome1", "hobby1"},
		[]string{"nome2", "sobrenome2", "hobby2"},
		[]string{"nome3", "sobrenome3", "hobby3"},
	}

	for _, v := range x {
		for _, vv := range v {
			fmt.Printf("%v ", vv)
		}
		fmt.Println()
	}

}
