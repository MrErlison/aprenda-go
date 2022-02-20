//Cap 5 ex 4
package main

import (
	"fmt"
)

func main() {
	x := 100
	fmt.Printf("%d, %#x, %b\n", x, x, x)
	x = x << 1
	fmt.Printf("%d, %#x, %b\n", x, x, x)
	x = x >> 1
	fmt.Printf("%d, %#x, %b", x, x, x)
}
