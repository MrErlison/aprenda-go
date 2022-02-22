// C7E4
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
