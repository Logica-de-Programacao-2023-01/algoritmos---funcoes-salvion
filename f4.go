package main

import "fmt"

//Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.

func segundo(num []int) int {
	maior := 0
	smaior := 0
	for _, valor := range num {
		if valor > maior {
			smaior = maior
			maior = valor
		}
		if valor < maior && valor > smaior {
			smaior = valor
		}
	}
	return smaior
}

func main() {
	num := []int{2, 3, 4, 5, 10, 1, 2}
	fmt.Print(segundo(num))
}
