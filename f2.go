package main

import (
	"fmt"
	"strings"
)

//Crie uma função que receba uma string e retorne a quantidade de vogais presentes.

func qntVogais(s string) int {
	vogais := []string{"a", "e", "i", "o", "u"}
	contador := 0

	for _, char := range strings.ToLower(s) {
		for _, vog := range vogais {
			if string(char) == vog {
				contador++
			}
		}
	}
	return contador
}

func main() {
	palavra := "Eeeiiitaaaa"
	fmt.Print(qntVogais(palavra))
}
