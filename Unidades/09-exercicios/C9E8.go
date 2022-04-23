// C9E8
// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.

package main

import "fmt"

func main() {

	x := map[string][]string{
		"nome1": []string{"nome1 hobby1", "nome1 hobby2", "nome1 hobby3"},
		"nome2": []string{"nome2 hobby1", "nome2 hobby2"},
		"nome3": []string{"nome3 hobby1", "nome3 hobby2", "nome3 hobby3"},
	}

	for i, v := range x {
		fmt.Println(i)
		for ii, vv := range v {
			fmt.Printf("\t%v %v ", ii, vv)
		}
		fmt.Println()
	}

}
