// C13E6
// - Crie e utilize uma função anônima.

package main

import "fmt"

func main() {
	func() {
		fmt.Println("Função anônima")
	}()
}
