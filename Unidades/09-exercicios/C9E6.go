// C9E6
// - Crie uma slice usando make que possa conter todos os estados do Brasil.
//     - Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
// - Demonstre o len e cap da slice.
// - Demonstre todos os valores da slice *sem utilizar range.*

package main

import "fmt"

func main() {

	x := make([]string, 26, 26)
	x = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia",
		"Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
		"Paraíba", "Paraná", "Pernambuco", "Piauí",
		"Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul",
		"Rondônia", "Roraima", "Santa Catarina", "São Paulo",
		"Sergipe", "Tocantins"}

	fmt.Println(len(x), cap(x))
	fmt.Println(x)
}
