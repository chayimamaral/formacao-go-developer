// usando pacote strings
package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := "A programação em Go é divertida"
	fmt.Println(strings.ToUpper(frase))

	// usando a função contains
	fmt.Println(strings.Contains(frase, "Go"))
	fmt.Println(strings.Contains(frase, "Python"))

	// usando a função replace
	novaFrase := strings.Replace(frase, "divertida", "incrível", 1)
	fmt.Println(novaFrase)

	// usando a função split
	palavras := strings.Split(frase, " ")
	fmt.Println(palavras)

	// usando a função trim
	texto := "   Olá, Mundo!   "
	textoTrimmed := strings.TrimSpace(texto)
	fmt.Printf("Antes: '%s'\nDepois: '%s'\n", texto, textoTrimmed)
}
