package main

import "fmt"

// Crie uma função que receba um slice de inteiros e um valor inteiro,
// e retorne a posição do primeiro elemento igual ao valor no slice. Caso não encontre,
// retorne -1.
func procura(num []int, n int) int {
	posicao := -1
	for i, valor := range num {
		if n == valor {
			posicao = i
			break
		}
	}
	if posicao == -1 {
		return -1
	} else {
		return posicao
	}
}

func main() {
	n := 0
	num := []int{2, 7, 3, 8, 4, 2, 1, 7, 5, 8, 9, 10}
	fmt.Print("Valor: ")
	fmt.Scan(&n)
	fmt.Print(procura(num, n))
}
