//recursão é a capacidade de uma função chamar a si mesma
//calcular o fatorial de um número

package main

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func main() {
	println(fatorial(2))
}
