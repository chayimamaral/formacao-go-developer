// servidor client -
// emitir uma solicitação de um client a um servidor http (../server/server.go
// e receber a resposta do servidor.

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// Emitir uma solicitação GET para o servidor
	response, err := http.Get("https://gobyexample.com")
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação:", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Status do servidor:", response.Status)

	// Ler a resposta do servidor
	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println("Erro ao ler a resposta:", err)
	// 	return
	// }

	scanner := bufio.NewScanner(response.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	// Imprimir a resposta recebida do servidor
	//fmt.Println("Resposta do servidor:", response.Body)
}
