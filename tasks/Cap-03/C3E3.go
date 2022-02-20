// golang day 2 ex 3

package main

import "fmt"

var x int = 42
var y string = "Jamed Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)

}
