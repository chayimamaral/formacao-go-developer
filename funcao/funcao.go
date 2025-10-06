package main

import (
	"fmt"
)

func calcularMedia(notas []float64) float64 {
	soma := 0.0
	for _, nota := range notas {
		soma += nota
	}
	return soma / float64(len(notas))
}

//programa que calcula a média de uma sala

func main() {
	lista := []float64{7.5, 8.0, 6.5, 9.0, 5.5}
	media := calcularMedia(lista)
	fmt.Printf("A média da sala é: %.2f\n", media)
}

