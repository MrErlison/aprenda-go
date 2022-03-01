// C13E2

package main

import "fmt"

func soma(v ...int) int {
	total := 0
	for _, x := range v {
		total += x
	}
	return total
}

func soma_si(v []int) int {
	total := 0
	for _, x := range v {
		total += x
	}
	return total
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(soma(x...))
	fmt.Println(soma_si(x))
}
