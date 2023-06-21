package main

import (
	"errors"
	"fmt"
)

// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice
// com os valores ordenados de forma crescente. Caso o slice esteja vazio, retorne um erro.
func ordenar(slice []int) ([]int, error) {
	if slice == nil {
		return nil, errors.New("slice vazio")
	}

	novo := slice

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if novo[j] < novo[i] {
				novo[i], novo[j] = novo[j], novo[i]
			}
		}
	}
	return novo, nil
}

func main() {
	slice := []int{2, 8, 1, 30, 62, 4, 22}
	novoSlice, err := ordenar(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Print(slice, novoSlice)
}
