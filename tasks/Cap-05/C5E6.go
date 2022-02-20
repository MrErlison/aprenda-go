package main

import (
	"fmt"
)

const (
	x = iota + 2022
	y
	z
	d
)

func main() {
	fmt.Println(x, y, z, d)
}
