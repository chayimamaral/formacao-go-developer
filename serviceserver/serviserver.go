// criar um servidor de servico que ira prestar algo ao usuario
package main

import (
	"fmt"
	"net/http"
)

func ola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ola!\n")
}

func cabecalho(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintln(w, "Cabecalho:\n")
	for nome, cabecalho := range req.Header {
		for _, c := range cabecalho {
			fmt.Fprintf(w, "%v: %v\n", nome, c)
		}
	}
}

func main() {
	http.HandleFunc("/ola", ola)             //quando alguem acessar a raiz do servidor, chama a funcao ola
	http.HandleFunc("/cabecalho", cabecalho) //quando alguem acessar a raiz do servidor, chama a funcao ola
	fmt.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", nil) //inicia o servidor na porta 8080
}
