// preciso criar um código em duas partes:
// desafio 1: um programa que trabalhe com o operador % e for.
// queremos criar um código que exiba todos os números entre 1 e 100 que são divisíveis por 3

// desafio 2: ao falar números, sempre que aparecer um múltiplo de 3, o participante deve falar "pin" 
// e nos múltiplos de 5 o participante deve falar 'pan'. Então o programa deve imprimir números de 1 a 100 e 
// nos casos cidados, não devem aparecer os números, mas sim, as palavras "pin" ou "pan".

package main

import "fmt"

func main() {

	//desafio 1
	fmt.Println("Números divisíveis por 3 entre 1 e 100:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}


	//desafio 2
	fmt.Println("\nNúmeros de 1 a 100 com 'pin' para múltiplos de 3 e 'pan' para múltiplos de 5:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i)
		} else if i%3 == 0 {
			fmt.Println("pin")
		} else if i%5 == 0 {
			fmt.Println("pan")
		} else {
			fmt.Println(i)
		}
	}
}
