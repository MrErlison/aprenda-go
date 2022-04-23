// C7E10
// - Anote (à mão) o resultado das expressões:
//     - fmt.Println(true && true)
//     - fmt.Println(true && false)
//     - fmt.Println(true || true)
//     - fmt.Println(true || false)
//     - fmt.Println(!true)

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
