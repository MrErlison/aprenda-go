// C7E10
package main

import "fmt"

func main() {
	x := true

	fmt.Println(x && x)
	fmt.Println(x && !x)
	fmt.Println(x || x)
	fmt.Println(x || !x)
	fmt.Println(!x)
}
