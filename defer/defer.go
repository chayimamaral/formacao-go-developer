//panic: exemplo de uso do panic e recover
//panic é usado para tratar erros críticos, que não podem ser tratados de outra forma
//recover é usado para recuperar o controle do programa após um panic

package main

import "fmt"

// func f2() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recuperando do panic:", r)
// 		}
// 	}()
// 	fmt.Println("f2")
// 	panic("Panic em f2")
// 	//fmt.Println("f2 depois do panic") --- IGNORE ---
// }

// func f3() {
// 	fmt.Println("f3")
// }

func main() {
	defer func() {
		x := recover()
		fmt.Println("Recuperando do panic na main:", x)
	}()

	panic("Panic na main")
}

// package main

// import "fmt"

// func dia1() {
// 	fmt.Println("Domingo")
// }

// func dia2() {
// 	fmt.Println("Segunda-feira")
// }

// func main() {
// 	defer dia2()
// 	dia1()
// }
