// C13E3

package main

import "fmt"

func main() {
	defer fmt.Println(0)

	for x := 1; x < 10; x++ {
		fmt.Println(x)
	}
}
