// go day 2 ex 5

package main

import "fmt"

type newtype int

var x newtype
var y int

func main() {

	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Printf("%v %T\n", x, x)

	y = int(x)
	fmt.Printf("%v %T", y, y)
}
