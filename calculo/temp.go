package main

import "fmt"

const ebulicaoF float64 = 212.0

func main() {

	tempC  := (ebulicaoF - 32) * 5 / 9
	fmt.Println("Constante ebulicaoF:", ebulicaoF)
	fmt.Printf("A temperatura de %.2f°F corresponde a %.2f°C\n", ebulicaoF, tempC)
}
