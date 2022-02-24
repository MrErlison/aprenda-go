// C9E1
package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	for idx, value := range array {
		fmt.Println(idx, value)
	}

	fmt.Printf("%T", array)
}
