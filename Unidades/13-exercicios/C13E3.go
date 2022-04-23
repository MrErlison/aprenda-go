// C13E3
// - Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.

package main

import "fmt"

func main() {
	defer fmt.Println(0)

	for x := 1; x < 10; x++ {
		fmt.Println(x)
	}
}
