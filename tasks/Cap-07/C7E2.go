// C7E1
package main

import "fmt"

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for y := 0; y < 3; y++ {
			fmt.Printf("\t%#U\n", x)
		}
	}
}
