// C13E10

package main

import "fmt"

func Sequencia() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	proximoInteiro := Sequencia()
	avancarInteiro := Sequencia()

	fmt.Println(proximoInteiro())
	fmt.Println(proximoInteiro())

	fmt.Println(avancarInteiro())

	fmt.Println(proximoInteiro())

	fmt.Println(avancarInteiro())

}
