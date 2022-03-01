// Closure

package main

import "fmt"

func main() {
	x := i()
	y := i()

	fmt.Println("x", x())
	fmt.Println("y", y())
	fmt.Println("x", x())
	fmt.Println("y", y())
	fmt.Println("x", x())
	fmt.Println("y", y())

}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
