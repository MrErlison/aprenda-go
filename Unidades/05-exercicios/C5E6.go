// C5E6
// - Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// - Demonstre estes valores.

package main

import (
	"fmt"
)

const (
	x = iota + 2022
	y
	z
	d
)

func main() {
	fmt.Println(x, y, z, d)
}
