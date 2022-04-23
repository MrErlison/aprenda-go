// C13E5
// - Crie um tipo "quadrado"
// - Crie um tipo "círculo"
// - Crie um método "área" para cada tipo que calcule e retorne a área da figura
//     - Área do círculo: 2 * π * raio
//     - Área do quadrado: lado * lado
// - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// - Crie um valor de tipo "quadrado"
// - Crie um valor de tipo "círculo"
// - Use a função "info" para demonstrar a área do "quadrado"
// - Use a função "info" para demonstrar a área do "círculo"

package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}
func (r retangulo) perimetro() float64 {
	return 2*r.largura + 2*r.altura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func medida(g geometria) {
	fmt.Println("Área", g.area())
	fmt.Println("Perímetro", g.perimetro())
}

func main() {
	fmt.Println("Retangulo")
	r := retangulo{largura: 3, altura: 4}
	medida(r)

	fmt.Println()

	fmt.Println("Circulo")
	c := circulo{raio: 5}
	medida(c)
}
