// // utilizar a função io.writestring

// package main

// import (
// 	"io/util"
// 	"log"
// 	"os"
// )

// func main() {
// 	_, err := io.WriteString(os.Stdout, "Escrevendo com WriteString\n")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// package main

// import (
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {
// 	data := []byte("Escrevendo com Write\n")
// 	err := os.WriteFile("output.txt", data, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	_, err = io.WriteString(file, "Adicionando mais texto com WriteString\n")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// abrir o arquivo output.txt e ler o conteúdo utilizando a função ioutil.ReadFile

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(os.Stdout, string(data))
	if err != nil {
		log.Fatal(err)
	}
}
