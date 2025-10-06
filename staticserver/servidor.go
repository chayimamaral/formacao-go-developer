// criar servidor estático que deverá ler a pagina local servidor.html
// o endereco do servidor deverá ser http://localhost:3000/servidor.html
// a pasta onde estará o arquivo servidor.html deverá ser chamada "staticserver"
// que é a mesma pasta onde se encontra o programa servidor.go
package main

import (
	"log"
	"net/http"
)

func main() {
	// fs := http.FileServer(http.Dir("."))
	// http.Handle("/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "servidor.html")
	})

	log.Println("Escutando na porta :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
