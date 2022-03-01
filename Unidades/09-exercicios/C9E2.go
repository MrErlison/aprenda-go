// C9E2
package main

import "fmt"

func main() {
	slice := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for idx, value := range slice {
		fmt.Println(idx, value)
	}

	fmt.Printf("%T", slice)
}
