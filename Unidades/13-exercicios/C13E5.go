// C13E5

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
