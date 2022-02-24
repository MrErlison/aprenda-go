// C9E3
package main

import "fmt"

func main() {
	slice := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("1-3", slice[:3])
	fmt.Println("5-10", slice[4:])
	fmt.Println("2-7", slice[1:7])
	fmt.Println("3-9", slice[2:len(slice)-1])
}
