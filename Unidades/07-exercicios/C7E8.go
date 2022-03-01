// C7E8
package main

import "fmt"

func main() {
	x := 30

	fmt.Print("Valor de x é ")
	switch {
	case x == 10:
		fmt.Println("dez")
	case x == 20:
		fmt.Println("vinte")
	case x == 30:
		fmt.Println("trinta")
	}
}
