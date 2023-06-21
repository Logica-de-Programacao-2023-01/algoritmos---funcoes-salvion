package main

import (
	"fmt"
)

// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice
// com apenas os números pares contidos no slice. Caso o slice esteja vazio, retorne um erro.
func pares(slice []int) []int {

	if len(slice) == 0 {
		return nil
	}

	pares := []int{}

	for _, valor := range slice {
		if valor%2 == 0 {
			pares = append(pares, valor)
		}
	}
	return pares
}

func main() {
	num := []int{2, 3, 4, 5, 6, 7, 8, 10}
	fmt.Print(pares(num))
}
