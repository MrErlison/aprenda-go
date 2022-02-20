// go day 2 ex 4

package main

import "fmt"

type newtype int

var x newtype

func main() {

	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Printf("%v %T\n", x, x)
}
