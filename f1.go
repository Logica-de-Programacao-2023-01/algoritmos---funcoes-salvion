package main

import "fmt"

//Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.

func media(num []int) float64 {
	soma := 0
	for _, x := range num {
		soma += x
	}
	media := float64(soma) / float64(len(num))
	return media
}

func main() {
	num := []int{10, 40, 60, 20}
	media := media(num)
	fmt.Printf("A média dos números é: %f", media)
}
