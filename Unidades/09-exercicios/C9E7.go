// C9E7
package main

import "fmt"

func main() {

	x := [][]string{
		[]string{"nome1", "sobrenome1", "hobby1"},
		[]string{"nome2", "sobrenome2", "hobby2"},
		[]string{"nome3", "sobrenome3", "hobby3"},
	}

	for _, v := range x {
		for _, vv := range v {
			fmt.Printf("%v ", vv)
		}
		fmt.Println()
	}

}
