// Exemplo de funções

package main

import "fmt"

func main() {
	basica()
	argumento("Oi, boa tarde!")
	fmt.Println(retorno("OI,", "boa noite"))
	total, quantidade := variatico(10, 10)
	fmt.Println(total, quantidade)

	si := []int{1, 2, 3, 4, 5}
	fmt.Println(enumerando(si...))

	//func like variable
	x := 2
	y := func(x int) int {

		return x * 2
	}
	fmt.Println(y(x))

	//anon
	func(z int) {
		fmt.Println(z * 2)
	}(x)

	z := inception()
	fmt.Println(z(x))

	//outra forma de chamar a funcao inception
	fmt.Println(inception()(10))

}

func basica() {
	fmt.Println("Oi, bom dia!")
}

func argumento(s string) {
	fmt.Println(s)
}

func retorno(x, y string) (string, string) {
	return x, y
}

func variatico(valor ...int) (int, int) {
	soma := 0
	for _, v := range valor {
		soma += v
	}
	return soma, len(valor)
}

func enumerando(valor ...int) int {
	soma := 0
	for _, v := range valor {
		soma += v
	}
	return soma
}

func inception() func(int) int {
	return func(x int) int {
		return x * 10
	}
}
