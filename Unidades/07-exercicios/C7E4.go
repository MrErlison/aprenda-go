// C7E4
// - Crie um loop utilizando a sintaxe: for {}
// - Utilize-o para demonstrar os anos desde que você nasceu.

package main

import "fmt"

func main() {
	ano_nascimento := 1900
	ano_final := 2090

	for {
		if ano_nascimento > ano_final {
			break
		}
		fmt.Println(ano_nascimento)
		ano_nascimento++
	}
}
