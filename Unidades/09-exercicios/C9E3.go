// C9E3
// - Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
//     - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
//     - Do quinto ao último item do slice (incluindo o último item!)
//     - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
//     - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
//     - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

package main

import "fmt"

func main() {
	slice := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("1-3", slice[:3])
	fmt.Println("5-10", slice[4:])
	fmt.Println("2-7", slice[1:7])
	fmt.Println("3-9", slice[2:len(slice)-1])
}
