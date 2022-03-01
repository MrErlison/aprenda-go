// C7E3
package main

import "fmt"

func main() {
	ano_nascimento := 1900
	ano_final := 2090

	for ano_nascimento <= ano_final {
		fmt.Println(ano_nascimento)
		ano_nascimento++
	}
}
