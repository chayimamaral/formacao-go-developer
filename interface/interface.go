package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	lado float64
}

func (c quadrado) area() float64 {
	return c.lado * c.lado
}

func (c quadrado) perimetro() float64 {
	return c.lado * 4
}

type circulo struct {
	radio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.radio
}

func main() {
	var g geometria

	g = quadrado{lado: 5}
	fmt.Printf("Quadrado - Area: %.2f, Perimetro: %.2f\n", g.area(), g.perimetro())

	g = circulo{radio: 3}
	fmt.Printf("Circulo - Area: %.2f, Perimetro: %.2f\n", g.area(), g.perimetro())
}
